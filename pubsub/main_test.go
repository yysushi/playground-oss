package main_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"cloud.google.com/go/pubsub"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

func TestPubSub(t *testing.T) {
	eg, ctx := errgroup.WithContext(context.Background())
	// pubslisher
	eg.Go(func() error {
		client, err := pubsub.NewClient(ctx, "project-id")
		if err != nil {
			return err
		}
		defer client.Close()
		fmt.Println("connected by publisher")
		it := client.Topics(ctx)
		for {
			t, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Printf("%s in publisher", err)
				return err
			}
			fmt.Println(t)
		}
		return nil
	})
	// provisioner
	eg.Go(func() error {
		client, err := pubsub.NewClient(ctx, "project-id")
		if err != nil {
			return err
		}
		defer client.Close()
		fmt.Println("connected by provisioner")
		_, err = client.CreateTopic(ctx, "testtopic")
		if err != nil {
			fmt.Printf("%s in provisioner", err)
			return err
		}
		fmt.Println("created topic")
		return nil
	})

	// subscriber
	eg.Go(func() error {
		client, err := pubsub.NewClient(ctx, "project-id")
		if err != nil {
			fmt.Printf("%s in subscriber", err)
			return err
		}
		defer client.Close()
		fmt.Println("connected by subscriber")
		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
