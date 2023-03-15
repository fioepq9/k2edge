package logic

import (
	"context"
	"k2edge/master/internal/types"
	"testing"
	"time"
)

func TestRegisterNode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := NewRegisterNodeLogic(ctx, &testSvcCtx)
	
	err := l.RegisterNode(&types.RegisterRequest{
		Name: "test1",
		Roles: []string{"master"},
		BaseURL: types.NodeURL{
			MasterURL: "43.138.169.60",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("register node success")
}

func TestDeleteNode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := NewDeleteNodeLogic(ctx, &testSvcCtx)
	
	err := l.DeleteNode(&types.DeleteRequest{
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("delete node success")
}

func TestNodeTop(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := NewNodeTopLogic(ctx, &testSvcCtx)
	
	info, err := l.NodeTop(&types.NodeTopRequest{
		Name: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
	t.Log("get node top info success")
}