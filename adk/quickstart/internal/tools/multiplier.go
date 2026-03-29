package tools

import (
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

func NewMultiplierTool() (tool.Tool, error) {
	multiply := func(_ tool.Context, args struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}) (map[string]any, error) {
		return map[string]any{"result": args.A * args.B}, nil
	}

	return functiontool.New(functiontool.Config{
		Name:        "multiplier",
		Description: "Multiplies two numbers and returns the result.",
	}, multiply)
}
