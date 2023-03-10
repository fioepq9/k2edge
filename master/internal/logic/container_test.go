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
				Name: "111",
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
		t.Fatal(err)
	}
	t.Log("create container success")
}

func TestDeleteContainerLogic(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	l := NewDeleteContainerLogic(ctx, &testSvcCtx)
	// l1 := NewCreateContainerLogic(ctx, &testSvcCtx)

	namespace := "system"
	containerName := "222"
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

func TestGetContainer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	l := NewGetContainerLogic(ctx, &testSvcCtx)
	
	container, err := l.GetContainer(&types.GetContainerRequest{
		Namespace: "system",
		Name: "111",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(container)
	t.Log("create container success")
}

func TestListContainer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	l := NewListContainerLogic(ctx, &testSvcCtx)
	
	containers, err := l.ListContainer(&types.ListContainerRequest{
		Namespace: "default",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(containers)
	t.Log("create container success")
}

func TestApplyContainer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	l := NewApplyContainerLogic(ctx, &testSvcCtx)
	
	err := l.ApplyContainer(&types.ApplyContainerRequest{
		Container: types.Container{
			Metadata: types.Metadata{
				Namespace: "default",
				Kind: "container",
				Name: "111",
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
		t.Fatal(err)
	}
	t.Log("apply container success")
}


func TestExecContainer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	l := NewExecContainerLogic(ctx, &testSvcCtx)
	
	err := l.ExecContainer(&types.ExecContainerRequest{
		Metadata: types.Metadata{
			Namespace: "default",
			Name: "111",
		},

		Config: types.ExecConfig{
			Cmd: []string{"ls"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("exec container success")
}