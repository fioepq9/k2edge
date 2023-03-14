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
		Name: "test",
		BaseURL: types.NodeURL{
			WorkerURL: "43.138.169.60",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("register node success")
}