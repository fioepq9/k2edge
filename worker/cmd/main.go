package main

import (
	"context"
	"io"
	"k2edge/worker/client"
	"os"
	"sync"
)

func main() {
	cli := client.NewClient(client.WithHost("localhost"), client.WithPort(8888))
	rw, err := cli.Container.Exec(context.Background(), client.ExecRequest{
		Container:    "f3",
		Tty:          true,
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          []string{`"/bin/bash"`},
	})
	if err != nil {
		panic(err)
	}
	defer rw.Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			if _, err := io.Copy(rw, os.Stdin); err != nil {
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if _, err := io.Copy(os.Stdout, rw); err != nil {
				break
			}
		}
	}()
	wg.Wait()
}
