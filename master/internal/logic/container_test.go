package logic

import (
	"context"
	"k2edge/master/internal/config"
	"k2edge/master/internal/svc"
	"k2edge/master/internal/types"
	"os"
	"testing"
	"time"
)

var testSvcCtx svc.ServiceContext

func TestMain(m *testing.M) {
	testSvcCtx = *svc.NewServiceContext(config.Config{
		Etcd: config.EtcdConf{
			Endpoints:   []string{"outlg.xyz:2379"},
			DialTimeout: 3,
		},
	})
	code := m.Run()
	os.Exit(code)
}

func TestCreatContainer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	l := NewCreateContainerLogic(ctx, &testSvcCtx)
	
	err := l.CreateContainer(&types.CreateContainerRequest{
		Container: types.Container{
			Metadata: types.Metadata{
				Namespace: "default",
				Kind: "container",
				Name: "333",
			},
			ContainerConfig: types.ContainerConfig{
				Image: "nginx",
				NodeName: "outlg",
			},
			ContainerStatus: types.ContainerStatus{
			},
		},
	})
	if err != nil {
		t.Log(err)
	}
	t.Log("create container success")
}

func TestDeleteContainerLogic(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	l := NewDeleteContainerLogic(ctx, &testSvcCtx)
	// l1 := NewCreateContainerLogic(ctx, &testSvcCtx)

	namespace := "default"
	containerName := "333"
	// err := l1.CreateContainer(&types.CreateContainerRequest{
	// 	Container: types.Container{
	// 		Metadata: types.Metadata{
	// 			Namespace: namespace,
	// 			Kind: "container",
	// 			Name: containerName,
	// 		},
	// 		ContainerConfig: types.ContainerConfig{
	// 			Image: "nginx",
	// 			NodeName: "outlg",
	// 		},
	// 		ContainerStatus: types.ContainerStatus{
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	t.Log(err)
	// }
	// t.Log("create container success")

	err := l.DeleteContainer(&types.DeleteContainerRequest{
		Namespace: namespace,
		Name: containerName,
		Timeout: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("delete container success")
}
