package main

import (
	"context"
	"io"
	"k2edge/worker/client"
	"os"
	"sync"
)

func main() {
	cli := client.NewClient("http://localhost:8888")
	// rw, err := cli.Container.Exec(context.Background(), client.ExecRequest{
	// 	Container:    "20ccbaf512101616125e981b2b623028c92a41dd1268982dba7e81ab1a41acb7",
	// 	Tty:          true,
	// 	AttachStdin:  true,
	// 	AttachStderr: true,
	// 	AttachStdout: true,
	// 	Cmd:          []string{`"/bin/bash"`},
	// })

	rw, err := cli.Container.Attach(context.Background(), client.AttachRequest{
		Container: "e6",
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
