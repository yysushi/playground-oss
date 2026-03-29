package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"

	// "google.golang.org/adk/model/gemini"
	"github.com/achetronic/adk-utils-go/genai/anthropic"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
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

	timeAgent, err := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       model,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(timeAgent),
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
