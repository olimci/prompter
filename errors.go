package prompter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ErrNoninteractive indicates that prompter could not acquire an interactive terminal.
var ErrNoninteractive = errors.New("prompter requires an interactive terminal")

func normalizeProgramError(err error) error {
	if err == nil {
		return nil
	}
	if isNoninteractiveError(err) {
		return fmt.Errorf("%w: %w", ErrNoninteractive, err)
	}
	return err
}

func isNoninteractiveError(err error) bool {
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		switch strings.ToLower(pathErr.Path) {
		case "/dev/tty", "conin$":
			return true
		}
	}

	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "could not open a new tty") ||
		strings.Contains(msg, "open /dev/tty") ||
		strings.Contains(msg, "open conin$")
}

func contextErr(ctx context.Context) error {
	if cause := context.Cause(ctx); cause != nil {
		return cause
	}
	return ctx.Err()
}

func resolveStartError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.Canceled) {
		if cause := context.Cause(ctx); cause != nil {
			if errors.Is(cause, context.Canceled) {
				return nil
			}
			return cause
		}
	}

	return err
}
