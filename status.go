package prompter

type StatusHandle struct {
	prompter *Prompter
	modal    *statusModal
}

func (p *Prompter) Status(message string) (*StatusHandle, error) {
	modal := newStatusModal(statusWorking, message, p.styles)
	if err := p.send(msgModal{modal: modal}); err != nil {
		return nil, err
	}
	return &StatusHandle{prompter: p, modal: modal}, nil
}

func (h *StatusHandle) Idle(message string) error {
	return h.prompter.send(msgStatusIdle{modal: h.modal, message: message})
}

func (h *StatusHandle) Working(message string) error {
	return h.prompter.send(msgStatusWorking{modal: h.modal, message: message})
}

func (h *StatusHandle) Progress(message string, percent float64) error {
	return h.prompter.send(msgStatusProgress{modal: h.modal, message: message, percent: percent})
}

func (h *StatusHandle) SetProgress(percent float64) error {
	return h.prompter.send(msgStatusProgressValue{modal: h.modal, percent: percent})
}

func (h *StatusHandle) Message(message string) error {
	return h.prompter.send(msgStatusMessage{modal: h.modal, message: message})
}

func (h *StatusHandle) Success(message string) error {
	return h.prompter.send(msgStatusSuccess{modal: h.modal, message: message})
}

func (h *StatusHandle) Error(message string) error {
	return h.prompter.send(msgStatusError{modal: h.modal, message: message})
}

func (h *StatusHandle) Clear() error {
	return h.prompter.send(msgStatusClear{modal: h.modal})
}
