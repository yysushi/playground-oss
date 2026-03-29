package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      interface{}     `json:"id,omitempty"`
}

type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	ID      interface{} `json:"id,omitempty"`
}

func main() {
	http.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		var result interface{}
		switch req.Method {
		case "initialize":
			result = map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"capabilities":    map[string]interface{}{"tools": map[string]interface{}{}},
				"serverInfo":      map[string]interface{}{"name": "simple-mcp", "version": "1.0.0"},
			}
		case "tools/list":
			result = map[string]interface{}{
				"tools": []map[string]interface{}{
					{
						"name":        "echo",
						"description": "Echo text",
						"inputSchema": map[string]interface{}{
							"type":       "object",
							"properties": map[string]interface{}{"text": map[string]string{"type": "string"}},
							"required":   []string{"text"},
						},
					},
					{
						"name":        "sum",
						"description": "Add two numbers",
						"inputSchema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"a": map[string]string{"type": "number"},
								"b": map[string]string{"type": "number"},
							},
							"required": []string{"a", "b"},
						},
					},
				},
			}
		case "tools/call":
			var params struct {
				Name      string                 `json:"name"`
				Arguments map[string]interface{} `json:"arguments"`
			}
			json.Unmarshal(req.Params, &params)

			if params.Name == "echo" {
				result = map[string]interface{}{
					"content": []map[string]string{{"type": "text", "text": params.Arguments["text"].(string)}},
				}
			} else if params.Name == "sum" {
				sum := params.Arguments["a"].(float64) + params.Arguments["b"].(float64)
				result = map[string]interface{}{
					"content": []map[string]string{{"type": "text", "text": fmt.Sprintf("%.0f", sum)}},
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{JSONRPC: "2.0", Result: result, ID: req.ID})
	})

	log.Println("MCP Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
