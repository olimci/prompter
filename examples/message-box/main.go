package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/olimci/prompter"
)

func boxedView(maxLines int) func(history []string, modal string) string {
	return func(history []string, modal string) string {
		lines := make([]string, 0, len(history)+1)
		lines = append(lines, history...)
		if modal != "" {
			lines = append(lines, modal)
		}
		if maxLines > 0 && len(lines) > maxLines {
			lines = lines[len(lines)-maxLines:]
		}

		width := 0
		for _, line := range lines {
			if w := lipgloss.Width(line); w > width {
				width = w
			}
		}

		top := "+" + strings.Repeat("-", width+2) + "+"
		var b strings.Builder
		b.WriteString(top)
		b.WriteString("\n")
		for _, line := range lines {
			pad := width - lipgloss.Width(line)
			if pad < 0 {
				pad = 0
			}
			b.WriteString("| ")
			b.WriteString(line)
			if pad > 0 {
				b.WriteString(strings.Repeat(" ", pad))
			}
			b.WriteString(" |")
			b.WriteString("\n")
		}
		b.WriteString(top)
		b.WriteString("\n")
		return b.String()
	}
}

func main() {
	ctx := context.Background()

	err := prompter.Start(func(p *prompter.Prompter) error {
		p.Log("Message box demo (type /done to exit)")

		box, err := p.MessageBox(
			prompter.WithInputPrompt("> "),
			prompter.WithInputSuggestions([]string{"/done", "/help", "/retry", "ping", "pong"}),
		)
		if err != nil {
			return err
		}

		for msg := range box.Messages() {
			if msg == "/done" {
				if err := box.Clear(); err != nil {
					return err
				}
				p.Log("goodbye")
				return nil
			}

			if msg == "/help" {
				p.Log("commands: /done, /help, /retry")
				continue
			}

			p.Logf("you said: %s", msg)
		}

		return nil
	},
		prompter.WithContext(ctx),
		prompter.WithMaxHistoryLen(10),
		prompter.WithViewFunc(boxedView(10)),
	)

	if err != nil {
		panic(fmt.Errorf("message box demo failed: %w", err))
	}
}
