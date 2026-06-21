package async

import (
	"context"
	"errors"
	"sync"
	"testing"
)

func TestTaskReturnsResult(t *testing.T) {
	task := New(context.Background(), func(context.Context) (string, error) {
		return "done", nil
	})

	got, err := task.Wait(context.Background())
	if err != nil {
		t.Fatalf("Wait() error = %v", err)
	}
	if got != "done" {
		t.Fatalf("Wait() = %q; want %q", got, "done")
	}
}

func TestTaskSupportsConcurrentWaiters(t *testing.T) {
	release := make(chan struct{})
	task := New(context.Background(), func(context.Context) (int, error) {
		<-release
		return 42, nil
	})

	const waiters = 8
	results := make([]int, waiters)
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, err := task.Wait(context.Background())
			if err != nil {
				t.Errorf("Wait() error = %v", err)
				return
			}
			results[i] = value
		}(i)
	}

	close(release)
	wg.Wait()
	for i, result := range results {
		if result != 42 {
			t.Fatalf("results[%d] = %d; want 42", i, result)
		}
	}
}

func TestWaitCanBeCanceled(t *testing.T) {
	block := make(chan struct{})
	task := New(context.Background(), func(context.Context) (string, error) {
		<-block
		return "done", nil
	})

	waitCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := task.Wait(waitCtx)
	close(block)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Wait() error = %v; want %v", err, context.Canceled)
	}
}
