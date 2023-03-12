package client

import (
	"context"
	"os"
	"testing"
	"time"
)

var testC *Client

func TestMain(m *testing.M) {
	testC = NewClient("http://localhost:8888")
	testC.EnableDumpAll()
	code := m.Run()
	os.Exit(code)
}

func TestContainer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sresp, err := testC.Containers().Status(ctx, ContainerStatusRequest{
		ID: "79f616e29b8b",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", sresp)
	err = testC.Containers().Exec(ctx, ExecRequest{
		Container: "79f616e29b8b",
		Config: ExecConfig{
			Tty:          true,
			AttachStdin:  true,
			AttachStderr: true,
			AttachStdout: true,
			Cmd:          []string{"/bin/bash"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
