package logic

import (
	"context"
	"k2edge/master/internal/types"
	"testing"
	"time"
)

func TestCreatNamespace(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := NewCreateNamespaceLogic(ctx, &testSvcCtx)
	
	err := l.CreateNamespace(&types.CreateNamespaceRequest{
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("create namespace success")
}

func TestDeleteNamespace(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := NewDeleteNamespaceLogic(ctx, &testSvcCtx)

	err := l.DeleteNamespace(&types.DeleteNamespaceRequest{
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("delete namespace success")
}

func TestGetNamespace(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	l := NewGetNamespaceLogic(ctx, &testSvcCtx)

	namespace, err := l.GetNamespace(&types.GetNamespaceRequest{
		Name: "system",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(namespace)
	t.Log("get namespace success")
}

func TestListNamespace(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	l := NewListNamespaceLogic(ctx, &testSvcCtx)

	namespace, err := l.ListNamespace(&types.ListNamespaceRequest{
		All: false,
	})
	if err != nil {
		t.Log(err)
	}
	t.Log(namespace)
	t.Log("get namespace success")
}