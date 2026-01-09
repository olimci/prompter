package prompter

import "fmt"

const keybindBufferSize = 16

type Keybind struct {
	Key         string
	Event       string
	Description string
}

type KeybindEvent struct {
	Key   string
	Event string
}

type KeybindHandle struct {
	prompter *Prompter
	modal    *keybindModal
	stream   <-chan KeybindEvent
}

func (p *Prompter) Keybinds(bindings []Keybind, opts ...KeybindOption) (*KeybindHandle, error) {
	if len(bindings) == 0 {
		return nil, fmt.Errorf("keybinds cannot be empty")
	}
	options := defaultKeybindOptions().apply(opts...)
	out := make(chan KeybindEvent, keybindBufferSize)
	modal := newKeybindModal(options, bindings, out, p.styles)
	if err := p.send(msgModal{modal: modal}); err != nil {
		return nil, err
	}
	return &KeybindHandle{
		prompter: p,
		modal:    modal,
		stream:   out,
	}, nil
}

func (h *KeybindHandle) Events() <-chan KeybindEvent {
	return h.stream
}

func (h *KeybindHandle) Set(bindings []Keybind) error {
	return h.prompter.send(msgKeybindSet{modal: h.modal, bindings: bindings})
}

func (h *KeybindHandle) Clear() error {
	return h.prompter.send(msgKeybindClear{modal: h.modal})
}
