package prompter

import "fmt"

type StatusKeybindHandle struct {
	prompter *Prompter
	modal    *statusKeybindModal
	status   *statusModal
	keybinds *keybindModal
	stream   <-chan KeybindEvent
}

func (p *Prompter) StatusKeybinds(message string, bindings []Keybind, opts ...KeybindOption) (*StatusKeybindHandle, error) {
	if len(bindings) == 0 {
		return nil, fmt.Errorf("keybinds cannot be empty")
	}
	options := defaultKeybindOptions().apply(opts...)
	out := make(chan KeybindEvent, keybindBufferSize)
	statusModal := newStatusModal(statusWorking, message, p.styles)
	keybindModal := newKeybindModal(options, bindings, out, p.styles)
	modal := newStatusKeybindModal(statusModal, keybindModal)
	if err := p.send(msgModal{modal: modal}); err != nil {
		return nil, err
	}
	return &StatusKeybindHandle{
		prompter: p,
		modal:    modal,
		status:   statusModal,
		keybinds: keybindModal,
		stream:   out,
	}, nil
}

func (h *StatusKeybindHandle) Events() <-chan KeybindEvent {
	return h.stream
}

func (h *StatusKeybindHandle) Set(bindings []Keybind) error {
	return h.prompter.send(msgKeybindSet{modal: h.keybinds, bindings: bindings})
}

func (h *StatusKeybindHandle) Idle(message string) error {
	return h.prompter.send(msgStatusIdle{modal: h.status, message: message})
}

func (h *StatusKeybindHandle) Working(message string) error {
	return h.prompter.send(msgStatusWorking{modal: h.status, message: message})
}

func (h *StatusKeybindHandle) Progress(message string, percent float64) error {
	return h.prompter.send(msgStatusProgress{modal: h.status, message: message, percent: percent})
}

func (h *StatusKeybindHandle) SetProgress(percent float64) error {
	return h.prompter.send(msgStatusProgressValue{modal: h.status, percent: percent})
}

func (h *StatusKeybindHandle) Message(message string) error {
	return h.prompter.send(msgStatusMessage{modal: h.status, message: message})
}

func (h *StatusKeybindHandle) Success(message string) error {
	return h.prompter.send(msgStatusSuccess{modal: h.status, message: message})
}

func (h *StatusKeybindHandle) Error(message string) error {
	return h.prompter.send(msgStatusError{modal: h.status, message: message})
}

func (h *StatusKeybindHandle) Clear() error {
	return h.prompter.send(msgStatusKeybindClear{modal: h.modal})
}
