package main_test

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var port = flag.Int("port", 0, "")

func runClient(addr string, name string) error {
	logger := slog.With("component", "client")
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("did not connect", "error", err)
		return err
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		logger.Error("could not greet", "error", err)
		return err
	}
	logger.Info("Greeting", "message", r.GetMessage())
	return nil
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	slog.With("component", "server").Info("Received", "name", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func runServer(ctx context.Context, port int) (string, error) {
	logger := slog.With("component", "server")
	var lc net.ListenConfig
	lis, err := lc.Listen(ctx, "tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("failed to listen", "error", err)
		return "", err
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	logger.Info("Listening", "address", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			logger.Error("failed to serve", "error", err)
		}
	}()
	return lis.Addr().String(), nil
}

func TestA(t *testing.T) {
	// port := flag.Int("port", 0, "")
	flag.Parse()
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	addr, err := runServer(ctx, *port)
	if err != nil {
		panic(err)
	}
	err = runClient(addr, "yysushi")
	if err != nil {
		panic(err)
	}
	cancel()
	<-ctx.Done()
}
