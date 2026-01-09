package main

import (
	"context"
	"fmt"

	"github.com/olimci/prompter"
)

func main() {
	ctx := context.Background()

	err := prompter.Start(func(ctx context.Context, p *prompter.Prompter) error {
		p.Log("Keybinds demo (press q to quit)")

		keys, err := p.Keybinds([]prompter.Keybind{
			{Key: "q", Event: "quit", Description: "quit"},
			{Key: "r", Event: "reload", Description: "reload"},
			{Key: "h", Event: "help", Description: "help"},
		}, prompter.WithKeybindPrompt("keys:"))
		if err != nil {
			return err
		}

		for ev := range keys.Events() {
			switch ev.Event {
			case "quit":
				_ = keys.Clear()
				p.Log("goodbye")
				return nil
			case "reload":
				p.Log("reload triggered")
			case "help":
				p.Log("available: q quit, r reload, h help")
			}
		}

		return nil
	}, prompter.WithContext(ctx))

	if err != nil {
		panic(fmt.Errorf("keybinds demo failed: %w", err))
	}
}
