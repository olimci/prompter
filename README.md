# prompter

A small Bubble Tea-based prompt helper for building interactive CLI flows.

## Install

```bash
go get github.com/olimci/prompter
```

## Quick usage

```go
package main

import (
	"context"

	"github.com/olimci/prompter"
)

func main() {
	ctx := context.Background()

	err := prompter.Start(func(ctx context.Context, p *prompter.Prompter) error {
		p.Log("Hello")

		name, err := p.AwaitInput(
			prompter.WithInputPrompt("Name: "),
		)
		if err != nil {
			return err
		}

		p.Logf("Hi %s", name)
		return nil
	}, prompter.WithContext(ctx))

	if err != nil {
		panic(err)
	}
}
```

## Examples

Run any of the examples:

```bash
go run ./examples/credit-card
go run ./examples/message-box
go run ./examples/status
go run ./examples/keybinds
go run ./examples/status-keybinds
go run ./examples/noninteractive
```
