package main_test

import (
	"context"
	"testing"
	"time"

	"github.com/ory/dockertest"
	"go.etcd.io/etcd/clientv3"
)

func TestMain(m *testing.M) {
	// etcd
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
	pool.Run()
	m.Run()

}

func TestSomething(t *testing.T) {
	// connect
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		panic(err)
	}
	defer cli.Close()
	c, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	err = cli.Sync(c)
	cancel()
	if err != nil {
		panic(err)
	}

}
