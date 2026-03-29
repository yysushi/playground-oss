package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/oauth2-proxy/mockoidc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	m, _ := mockoidc.Run()
	defer func() { _ = m.Server.Shutdown(ctx) }()

	cfg := m.Config()
	fmt.Printf("%#v\n", cfg)

	select {
	case <-time.After(time.Second*30):
		fmt.Println("done")
	case <-ctx.Done():
		stop()
		fmt.Println("canceled")
	}
}
