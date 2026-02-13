package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/olimci/prompter"
)

func main() {
	ctx := context.Background()

	err := prompter.Start(func(ctx context.Context, p *prompter.Prompter) error {
		p.Log("Interactive terminal detected")
		return nil
	}, prompter.WithContext(ctx))
	if err == nil {
		return
	}

	if errors.Is(err, prompter.ErrNoninteractive) {
		fmt.Println("Noninteractive terminal detected; falling back to plain output mode.")
		fmt.Println("Hello from fallback mode.")
		return
	}

	panic(fmt.Errorf("noninteractive demo failed: %w", err))
}
