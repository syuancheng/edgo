package async

import "context"

type result[T any] struct {
	value T
	err   error
}

// Task represents one function that starts immediately and publishes one result.
// Closing done establishes the synchronization needed for concurrent Wait calls.
type Task[T any] struct {
	done   chan struct{}
	result result[T]
}

func New[T any](ctx context.Context, fn func(context.Context) (T, error)) *Task[T] {
	task := &Task[T]{done: make(chan struct{})}
	go func() {
		defer close(task.done)
		task.result.value, task.result.err = fn(ctx)
	}()
	return task
}

// Wait waits for the task or for waitCtx to be canceled. Multiple goroutines
// may call Wait concurrently and receive the same completed result.
func (t *Task[T]) Wait(waitCtx context.Context) (T, error) {
	select {
	case <-t.done:
		return t.result.value, t.result.err
	case <-waitCtx.Done():
		var zero T
		return zero, waitCtx.Err()
	}
}
