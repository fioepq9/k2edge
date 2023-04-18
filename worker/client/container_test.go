package client

import (
	"context"
	"k2edge/auth"
	"os"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var TestClient *Client

func TestMain(m *testing.M) {
	e, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"outlg.xyz:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	TestClient = NewClient(
		"http://localhost:8888",
		WithSecret("worker-dev-2", "D19DD0CC-70EB-49A6-98FA-82BA2404ED01"),
		WithAuther(auth.NewEtcdAuther("worker-dev-2", e)),
	)
	code := m.Run()
	e.Close()
	os.Exit(code)
}

func TestContainerList(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	resp, err := TestClient.Container.List(ctx, ListContainersRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
