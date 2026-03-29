package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"

	"github.com/achetronic/adk-utils-go/genai/anthropic"
	"google.golang.org/adk/tool"

	internaltools "sample/internal/tools"
)

func main() {
	ctx := context.Background()

	model := anthropic.New(anthropic.Config{
		APIKey:    os.Getenv("ANTHROPIC_API_KEY"),
		BaseURL:   os.Getenv("ANTHROPIC_BASE_URL"),
		ModelName: os.Getenv("ANTHROPIC_DEFAULT_OPUS_MODEL"),
		// MaxOutputTokens:      getEnvInt("MAX_OUTPUT_TOKENS", 0),
		// ThinkingBudgetTokens: getEnvInt("THINKING_BUDGET_TOKENS", 0),
	})

	multiplier, err := internaltools.NewMultiplierTool()
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	calculatorAgent, err := llmagent.New(llmagent.Config{
		Name:        "hello_calculator_agent",
		Model:       model,
		Description: "Tells an expression to calculate",
		Instruction: "You are a helpful assistant that calculates something.",
		Tools:       []tool.Tool{multiplier},
	})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(calculatorAgent),
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
