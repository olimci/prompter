package main

import (
	"context"
	"fmt"
	"time"

	"github.com/olimci/prompter"
)

func main() {
	ctx := context.Background()

	err := prompter.Start(func(p *prompter.Prompter) error {
		p.Log("Status modal demo")

		status, err := p.Status("Booting up")
		if err != nil {
			return err
		}

		time.Sleep(400 * time.Millisecond)
		if err := status.Working("Preparing work queue"); err != nil {
			return err
		}

		time.Sleep(400 * time.Millisecond)
		if err := status.Progress("Downloading assets", 0.0); err != nil {
			return err
		}

		for i := 1; i <= 10; i++ {
			time.Sleep(200 * time.Millisecond)
			if err := status.SetProgress(float64(i) / 10.0); err != nil {
				return err
			}
		}

		if err := status.Message("Finalizing"); err != nil {
			return err
		}
		time.Sleep(300 * time.Millisecond)

		if err := status.Clear(); err != nil {
			return err
		}

		p.Log("Done")
		return nil
	},
		prompter.WithContext(ctx),
	)

	if err != nil {
		panic(fmt.Errorf("status demo failed: %w", err))
	}
}
