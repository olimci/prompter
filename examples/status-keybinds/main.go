package main

import (
	"context"
	"fmt"
	"time"

	"github.com/olimci/prompter"
)

func main() {
	ctx := context.Background()

	err := prompter.Start(func(ctx context.Context, p *prompter.Prompter) error {
		p.Log("Status + keybinds demo (press q to quit)")

		status, err := p.StatusKeybinds("Booting up", []prompter.Keybind{
			{Key: "q", Event: "quit", Description: "quit"},
			{Key: "h", Event: "help", Description: "help"},
		}, prompter.WithKeybindPrompt("keys:"))
		if err != nil {
			return err
		}

		stop := make(chan struct{})
		done := make(chan struct{})
		go func() {
			defer close(done)

			sleep := func(d time.Duration) bool {
				select {
				case <-stop:
					return false
				case <-time.After(d):
					return true
				}
			}

			if !sleep(400 * time.Millisecond) {
				return
			}
			if err := status.Working("Preparing work queue"); err != nil {
				return
			}

			if !sleep(400 * time.Millisecond) {
				return
			}
			if err := status.Progress("Downloading assets", 0.0); err != nil {
				return
			}

			for i := 1; i <= 10; i++ {
				if !sleep(200 * time.Millisecond) {
					return
				}
				if err := status.SetProgress(float64(i) / 10.0); err != nil {
					return
				}
			}

			if err := status.Message("Finalizing"); err != nil {
				return
			}
			if !sleep(300 * time.Millisecond) {
				return
			}

			if err := status.Success("Done"); err != nil {
				return
			}
			if !sleep(250 * time.Millisecond) {
				return
			}

			_ = status.Clear()
		}()

		for {
			select {
			case <-done:
				return nil
			case ev, ok := <-status.Events():
				if !ok {
					return nil
				}
				switch ev.Event {
				case "quit":
					close(stop)
					_ = status.Clear()
					p.Log("goodbye")
					return nil
				case "help":
					p.Log("available: q quit, h help")
				}
			}
		}
	}, prompter.WithContext(ctx))

	if err != nil {
		panic(fmt.Errorf("status-keybinds demo failed: %w", err))
	}
}
