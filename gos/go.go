package gos

import "context"

// 协程池
func NewGoroutinePool(ctx context.Context, concurrent int) *GoroutinePool {
	newCtx, newCancel := context.WithCancel(ctx)
	gp := &GoroutinePool{
		ctx:    newCtx,
		cancel: newCancel,
		taskC:  make(chan Task, concurrent),
	}
	for i := 0; i < concurrent; i++ {
		go gp.worker()
	}
	return gp
}

type Task func(ctx context.Context)

type GoroutinePool struct {
	ctx    context.Context
	cancel context.CancelFunc
	taskC  chan Task
}

func (g *GoroutinePool) worker() {
	for {
		select {
		case <-g.ctx.Done():
			return
		case task := <-g.taskC:
			task(g.ctx)
		}
	}
}

func (g *GoroutinePool) Close() {
	if g.cancel != nil {
		g.cancel()
	}
}

func (g *GoroutinePool) BlockSubmit(ctx context.Context, task Task) {
	select {
	case <-g.ctx.Done():
	case <-ctx.Done():
	case g.taskC <- task:
	}
}

// errGroup
