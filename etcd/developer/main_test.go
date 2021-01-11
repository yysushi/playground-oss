package main_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
	etcd "go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

func newEnv() (*dockertest.Pool, *dockertest.Resource) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	options := &dockertest.RunOptions{
		Repository: "quay.io/coreos/etcd",
		Tag:        "v3.4.13",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"2379/tcp": []docker.PortBinding{{HostPort: "2379"}},
		},
		Cmd: []string{
			"/usr/local/bin/etcd",
			"--advertise-client-urls", "http://127.0.0.1:2379",
			"--listen-client-urls", "http://0.0.0.0:2379",
		},
	}
	resource, err := pool.RunWithOptions(options)
	if err != nil {
		panic(fmt.Sprintf("Failed to run etcd %s", err))
	}
	return pool, resource
}

func newClient() *etcd.Client {
	cli, err := etcd.New(etcd.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to new client: %s", err))
	}
	return cli
}

func TestConnection(t *testing.T) {
	// given
	pool, resource := newEnv()
	defer pool.Purge(resource)
	// when
	cli := newClient()
	defer cli.Close()
	// then
	connInfo := cli.ActiveConnection().Target()
	rx := regexp.MustCompile("endpoint://client-[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|a|b][a-f0-9]{3}-[a-f0-9]{12}/localhost:2379")
	assert.Regexp(t, rx, connInfo)
}

func TestGet(t *testing.T) {
	// given
	ctx := context.Background()
	pool, resource := newEnv()
	defer pool.Purge(resource)
	cli := newClient()
	defer cli.Close()
	defer cli.Delete(ctx, "", etcd.WithPrefix())
	// given
	cli.Put(ctx, "/key", "value")
	// when
	resp, err := cli.Get(context.TODO(), "/", etcd.WithPrefix())
	// then
	if err != nil {
		t.Fatalf("Failed to get: %s", err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s\n", ev.Key)
	}
}
