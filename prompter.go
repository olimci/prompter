package prompter

import (
	"context"
	"fmt"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/olimci/prompter/promise"
)

func newPrompter(o *options) *Prompter {
	ctx, cancel := context.WithCancelCause(o.context)
	if o.styles == nil {
		o.styles = DefaultStyles()
	}
	return &Prompter{
		ctx:     ctx,
		cancel:  cancel,
		styles:  o.styles,
		program: tea.NewProgram(newModel(o), tea.WithContext(ctx)),
	}
}

type Prompter struct {
	ctx     context.Context
	cancel  context.CancelCauseFunc
	program *tea.Program
	styles  *Styles
}

func (p *Prompter) send(msg tea.Msg) error {
	select {
	case <-p.ctx.Done():
		return contextErr(p.ctx)
	default:
	}

	p.program.Send(msg)

	return nil
}

func (p *Prompter) Log(msg string) {
	p.send(msgLog{message: msg}) // likely safe to discard (cope)
}

func (p *Prompter) Logf(format string, args ...any) {
	p.send(msgLog{message: fmt.Sprintf(format, args...)})
}

func (p *Prompter) Clear() {
	p.send(msgClearHistory{})
}

func (p *Prompter) Confirm(prompt string) (promise.Promise[bool], error) {
	pro, res := promise.New[bool]()
	if err := p.send(msgModal{newConfirmModal(prompt, true, res, p.styles)}); err != nil {
		return promise.Promise[bool]{}, err
	}
	return pro, nil
}

func (p *Prompter) AwaitConfirm(prompt string) (bool, error) {
	pro, res := promise.New[bool]()
	if err := p.send(msgModal{newConfirmModal(prompt, true, res, p.styles)}); err != nil {
		return false, err
	}
	return pro.Await(p.ctx)
}

func (p *Prompter) Input(opts ...InputOption) (promise.Promise[string], error) {
	pro, res := promise.New[string]()
	options := defaultInputOptions().apply(opts...)
	if err := p.send(msgModal{newInputModal(options, res, p.styles)}); err != nil {
		return promise.Promise[string]{}, err
	}

	return pro, nil
}

func (p *Prompter) AwaitInput(opts ...InputOption) (string, error) {
	pro, res := promise.New[string]()
	options := defaultInputOptions().apply(opts...)
	if err := p.send(msgModal{newInputModal(options, res, p.styles)}); err != nil {
		return "", err
	}
	return pro.Await(p.ctx)
}

func (p *Prompter) Select(prompt string, options []string) (promise.Promise[string], error) {
	if len(options) == 0 {
		return promise.Promise[string]{}, fmt.Errorf("select options cannot be empty")
	}
	pro, res := promise.New[string]()
	if err := p.send(msgModal{newSelectModal(prompt, options, 0, res, p.styles)}); err != nil {
		return promise.Promise[string]{}, err
	}
	return pro, nil
}

func (p *Prompter) AwaitSelect(prompt string, options []string) (string, error) {
	if len(options) == 0 {
		return "", fmt.Errorf("select options cannot be empty")
	}
	pro, res := promise.New[string]()
	if err := p.send(msgModal{newSelectModal(prompt, options, 0, res, p.styles)}); err != nil {
		return "", err
	}
	return pro.Await(p.ctx)
}

func (p *Prompter) SelectDefault(prompt string, options []string, defaultValue string) (promise.Promise[string], error) {
	if len(options) == 0 {
		return promise.Promise[string]{}, fmt.Errorf("select options cannot be empty")
	}

	index := slices.Index(options, defaultValue)
	if index == -1 {
		return promise.Promise[string]{}, fmt.Errorf("default value not found in options")
	}
	pro, res := promise.New[string]()
	if err := p.send(msgModal{newSelectModal(prompt, options, index, res, p.styles)}); err != nil {
		return promise.Promise[string]{}, err
	}
	return pro, nil
}

func (p *Prompter) AwaitSelectDefault(prompt string, options []string, defaultValue string) (string, error) {
	if len(options) == 0 {
		return "", fmt.Errorf("select options cannot be empty")
	}

	index := slices.Index(options, defaultValue)
	if index == -1 {
		return "", fmt.Errorf("default value not found in options")
	}
	pro, res := promise.New[string]()
	if err := p.send(msgModal{newSelectModal(prompt, options, index, res, p.styles)}); err != nil {
		return "", err
	}
	return pro.Await(p.ctx)
}

func Start(f func(ctx context.Context, p *Prompter) error, opts ...Option) error {
	o := defaultOptions().apply(opts...)
	p := newPrompter(o)
	defer func() {
		p.program.Quit()
		p.program.Wait()
	}()

	go func() {
		if _, err := p.program.Run(); err != nil {
			p.cancel(normalizeProgramError(err))
		} else {
			p.cancel(nil)
		}
	}()

	return resolveStartError(p.ctx, f(p.ctx, p))
}
