package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type EchoParams struct {
	Text string `json:"text" jsonschema:"Text to echo back"`
}

type SumParams struct {
	A float64 `json:"a" jsonschema:"First number"`
	B float64 `json:"b" jsonschema:"Second number"`
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "simple-mcp",
		Version: "1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "echo",
		Description: "Echo text",
	}, func(ctx context.Context, req *mcp.CallToolRequest, params *EchoParams) (*mcp.CallToolResult, any, error) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: params.Text},
			},
		}, nil, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "sum",
		Description: "Add two numbers",
	}, func(ctx context.Context, req *mcp.CallToolRequest, params *SumParams) (*mcp.CallToolResult, any, error) {
		result := params.A + params.B
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("%.0f", result)},
			},
		}, nil, nil
	})

	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		return server
	}, nil)

	log.Println("MCP Server listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
