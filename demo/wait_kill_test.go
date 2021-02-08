package demo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWaitKill(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel() // 防止任务比超时时间短导致资源未释放
	// 启动协程
	go task(ctx)
	// 主协程需要等待，否则直接退出
	time.Sleep(time.Second * 3)
}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	// 真正的任务协程
	go func() {
		// 模拟两秒耗时任务
		time.Sleep(time.Second * 2)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
