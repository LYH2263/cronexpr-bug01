package cronexpr

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LYH2263/go-cronexpr/internal/field"
	"github.com/LYH2263/go-cronexpr/internal/tz"
)

// Next returns the first fire time strictly after from in expr timezone.
func (e *Expr) Next(from time.Time) time.Time {
	t, err := e.NextWithContext(context.Background(), from)
	if err != nil {
		return time.Time{}
	}
	return t
}

// NextWithContext computes next fire time honoring cancellation.
func (e *Expr) NextWithContext(ctx context.Context, from time.Time) (time.Time, error) {
	if e == nil || e.tree == nil {
		return time.Time{}, ErrNilExpr
	}
	if err := ctx.Err(); err != nil {
		return time.Time{}, fmt.Errorf("%w: %v", ErrCanceled, err)
	}
	layer := tz.NewLayer(e.loc)
	local, err := layer.ConvertIn(ctx, from)
	if err != nil {
		return time.Time{}, wrapCancel(err)
	}
	refs := make([]field.FieldRef, len(e.tree.Fields))
	for i, f := range e.tree.Fields {
		refs[i] = field.FieldRef{Kind: f.Kind, Bitmap: f.Bitmap}
	}
	ev := field.NewEvaluator(refs, e.withSecond)
	t, ok, err := ev.NextAfter(ctx, local)
	if err != nil {
		return time.Time{}, wrapCancel(err)
	}
	if !ok {
		return time.Time{}, ErrNoMatch
	}
	t = layer.AdjustDST(t)
	return layer.ToUTC(t), nil
}

func wrapCancel(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return fmt.Errorf("%w: %v", ErrCanceled, err)
	}
	return err
}

// NextN returns up to n upcoming fire times after from.
func (e *Expr) NextN(from time.Time, n int) ([]time.Time, error) {
	return e.NextNWithContext(context.Background(), from, n)
}

// NextNWithContext returns up to n fire times, discarding partial results on error.
func (e *Expr) NextNWithContext(ctx context.Context, from time.Time, n int) ([]time.Time, error) {
	if n <= 0 {
		return nil, nil
	}
	out := make([]time.Time, 0, n)
	cur := from
	for len(out) < n {
		nxt, err := e.NextWithContext(ctx, cur)
		if err != nil {
			return nil, err
		}
		if nxt.IsZero() {
			break
		}
		out = append(out, nxt)
		cur = nxt
	}
	return out, nil
}
