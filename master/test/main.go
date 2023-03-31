package main

import (
	"context"
	"io"
	"k2edge/master/test/client"
	"os"
	"time"
)

func main() {
	testExecAttach()
}

func testExecAttach() {
	cli := client.NewClient("http://localhost:8080")
	rw, err := cli.Container.Exec(context.Background(), client.ExecContainerRequest{
		Namespace:    "default",
		Name:         "ccc",
		Tty:          true,
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          []string{`"\"/bin/sh\""`},
	})

	// rw, err := cli.Container.Attach(context.Background(), client.AttachContainerRequest{
	// 	Namespace: "default",
	// 	Name: "ccc",
	// })
	if err != nil {
		panic(err)
	}
	defer rw.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		for {
			if _, err := io.Copy(rw, os.Stdin); err != nil {
				break
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}()
	go func() {
		defer cancel()
		for {
			if _, err := io.Copy(os.Stdout, rw); err != nil {
				break
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}()
	<-ctx.Done()
}

func testLogs() {
	cli := client.NewClient("http://localhost:8080")

	rw, err := cli.Container.Logs(context.Background(), client.LogsContainerRequest{
		Namespace: "default",
		Name: "111",
		Follow: true,
	})
	if err != nil {
		panic(err)
	}
	defer rw.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		for {
			if _, err := io.Copy(os.Stdout, rw); err != nil {
				break
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}()
	<-ctx.Done()
	time.Sleep(time.Second)
}