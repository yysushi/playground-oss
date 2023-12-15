package main_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/status"
)

var knownErrMsgs = []string{
	"Topic already exists",
	"Subscription already exists",
}

func isKnownError(err error) bool {
	s, ok := status.FromError(err)
	for _, errMsg := range knownErrMsgs {
		if ok && s.Message() == errMsg {
			return true
		}
	}
	return false
}

func setUp(t *testing.T, ctx context.Context) (client *pubsub.Client) {
	var err error
	client, err = pubsub.NewClient(ctx, "testproject")
	if err != nil {
		t.Fatal(err)
	}
	var topic *pubsub.Topic
	topic, err = client.CreateTopic(ctx, "testtopic")
	if err != nil {
		if !isKnownError(err) {
			t.Fatal(err)
		}
		topic = client.Topic("testtopic")
	}
	_, err = client.CreateSubscription(ctx, "testsubscription", pubsub.SubscriptionConfig{
		Topic:            topic,
		AckDeadline:      10 * time.Second,
		ExpirationPolicy: time.Duration(0),
	})
	if err != nil && !isKnownError(err) {
		t.Fatal(err)
	}
	return
}

func tearDown(t *testing.T, ctx context.Context, client *pubsub.Client) {
	client.Close()
}

func TestPubSub_topics(t *testing.T) {
	ctx := context.Background()
	client := setUp(t, ctx)
	defer tearDown(t, ctx, client)

	it := client.Topics(ctx)
	topics := make([]*pubsub.Topic, 0)
	for {
		topic, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			} else {
				t.Fatal(err)
			}
		}
		topics = append(topics, topic)
	}
	assert.Len(t, topics, 1)
	assert.Equal(t, "testtopic", topics[0].ID())
}

type OrderResult int64

const (
	InOrder OrderResult = iota
	UnOrder
)

// Reception order is not the same as sent
// (use SubscriptionConfig.EnableMessageOrdering if needed)
// https://github.com/googleapis/google-cloud-go/blob/pubsub/v1.33.0/pubsub/subscription.go#L533
func TestPubSub_order(t *testing.T) {
	var err error
	ctx := context.Background()
	client := setUp(t, ctx)
	defer tearDown(t, ctx, client)

	met := make(map[OrderResult]struct{})
	var i = 0
	for ; i < 100; i++ {
		result := func() OrderResult {
			topic := client.Topic("testtopic")
			res := topic.Publish(ctx, &pubsub.Message{Data: []byte("1")})
			_, err = res.Get(ctx)
			if err != nil {
				log.Fatal(err)
			}
			res = topic.Publish(ctx, &pubsub.Message{Data: []byte("2")})
			_, err = res.Get(ctx)
			if err != nil {
				log.Fatal(err)
			}

			cctx, cancel := context.WithCancel(ctx)
			msgs := make([]string, 0)
			receivedMsgs := make(chan string)
			done := make(chan error)
			subscription := client.Subscription("testsubscription")
			go func() {
				done <- subscription.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
					receivedMsgs <- string(msg.Data)
					msg.Ack()
				})
			}()
			msgs = append(msgs, <-receivedMsgs)
			msgs = append(msgs, <-receivedMsgs)
			cancel()
			if err = <-done; err != nil && !errors.Is(err, context.Canceled) {
				t.Fatal(err)
			}
			assert.ElementsMatch(t, []string{"1", "2"}, msgs)
			if msgs[0] == "1" {
				return InOrder
			}
			return UnOrder
		}()
		met[result] = struct{}{}
		if l := len(met); l == 1 {
			fmt.Println(i, result)
		} else if l == 2 {
			fmt.Println(i, result)
			return
		}
	}
	t.Fatal("reached to max loop")
}

// Nack affects only one message
// (the nacked message is only resent)
func TestPubSub_nack(t *testing.T) {
	var err error
	ctx := context.Background()
	client := setUp(t, ctx)
	defer tearDown(t, ctx, client)

	topic := client.Topic("testtopic")
	res := topic.Publish(ctx, &pubsub.Message{Data: []byte("1")})
	_, err = res.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	res = topic.Publish(ctx, &pubsub.Message{Data: []byte("2")})
	_, err = res.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan error)
	cctx, cancel := context.WithCancel(ctx)
	receivedMsgs := make(chan *pubsub.Message)
	ret := make(chan func())
	subscription := client.Subscription("testsubscription")
	go func() {
		done <- subscription.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
			receivedMsgs <- msg
			f := <-ret
			f()
		})
	}()

	firstMsg := <-receivedMsgs
	ret <- firstMsg.Nack
	secondMsg := <-receivedMsgs
	ret <- secondMsg.Ack
	thirdMsg := <-receivedMsgs
	ret <- thirdMsg.Ack

	cancel()
	if err = <-done; err != nil && !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	assert.Equal(t, firstMsg.Data, thirdMsg.Data)
	assert.NotEqual(t, firstMsg.Data, secondMsg.Data)
}
