package prompter

import (
	"context"
	"errors"
	"os"
	"testing"
)

func TestNormalizeProgramErrorNoninteractive(t *testing.T) {
	pathErr := &os.PathError{
		Op:   "open",
		Path: "/dev/tty",
		Err:  os.ErrNotExist,
	}

	err := normalizeProgramError(pathErr)

	if !errors.Is(err, ErrNoninteractive) {
		t.Fatalf("expected ErrNoninteractive, got %v", err)
	}
	if !errors.Is(err, pathErr) {
		t.Fatalf("expected wrapped path error, got %v", err)
	}
}

func TestResolveStartErrorUsesCancelCause(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(normalizeProgramError(&os.PathError{
		Op:   "open",
		Path: "/dev/tty",
		Err:  os.ErrNotExist,
	}))

	err := resolveStartError(ctx, context.Canceled)

	if !errors.Is(err, ErrNoninteractive) {
		t.Fatalf("expected ErrNoninteractive, got %v", err)
	}
}

func TestResolveStartErrorGracefulCancelReturnsNil(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(nil)

	err := resolveStartError(ctx, context.Canceled)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestSendReturnsCancelCause(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(normalizeProgramError(&os.PathError{
		Op:   "open",
		Path: "/dev/tty",
		Err:  os.ErrNotExist,
	}))
	p := &Prompter{ctx: ctx}

	err := p.send(msgLog{message: "hello"})
	if !errors.Is(err, ErrNoninteractive) {
		t.Fatalf("expected ErrNoninteractive, got %v", err)
	}
}
