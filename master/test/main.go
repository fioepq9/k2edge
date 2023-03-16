package main

import (
	"context"
	"io"
	"k2edge/master/test/client"
	"os"
	"sync"
)

func main() {
	cli := client.NewClient("http://localhost:8080")
	// rw, err := cli.Container.Exec(context.Background(), client.ExecContainerRequest{
	// 	Namespace:    "default",
	// 	Name:         "333",
	// 	Tty:          true,
	// 	AttachStdin:  true,
	// 	AttachStderr: true,
	// 	AttachStdout: true,
	// 	Cmd:          []string{`"\"/bin/bash\""`},
	// })

	rw, err := cli.Container.Attach(context.Background(), client.AttachContainerRequest{
		Namespace: "default",
		Name: "ccc",
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
