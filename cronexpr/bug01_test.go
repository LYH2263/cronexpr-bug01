package cronexpr

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestBug01_NextScanHonorsCancel(t *testing.T) {
	e, err := Parse("0 0 31 2 *")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel()
	}()
	done := make(chan error, 1)
	go func() {
		_, err := e.NextWithContext(ctx, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
		done <- err
	}()
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("expected cancel error")
		}
		if !errors.Is(err, ErrCanceled) && !errors.Is(err, context.Canceled) {
			t.Fatalf("want canceled, got %v", err)
		}
	case <-time.After(300 * time.Millisecond):
		t.Fatal("scan ignored cancel")
	}
}
