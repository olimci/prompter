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
	_ = prompter.Start(func(p *prompter.Prompter) error {
		p.Log("Hello")

		namePro, err := p.Input(
			prompter.WithInputPrompt("Name: "),
		)
		if err != nil {
			return err
		}

		name, err := namePro.Await(p.Ctx)
		if err != nil {
			return err
		}

		p.Logf("Hi %s", name)
		return nil
	}, prompter.WithContext(context.Background()))
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
```
