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
		Name: "ljf",
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
		Name: "ljf",
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

func TestCordon(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := NewCordonLogic(ctx, &testSvcCtx)
	
	err := l.Cordon(&types.CordonRequest{
		Name: "outlg",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(err)
	t.Log("cordon node success")
}

func TestUncordon(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	l := NewUncordonLogic(ctx, &testSvcCtx)
	
	err := l.Uncordon(&types.UncordonRequest{
		Name: "outlg",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(err)
	t.Log("cordon node success")
}