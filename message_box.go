package prompter

const messageBoxBufferSize = 16

type MessageBoxHandle struct {
	prompter *Prompter
	modal    *messageBoxModal
	stream   <-chan string
}

func (p *Prompter) MessageBox(opts ...InputOption) (*MessageBoxHandle, error) {
	options := defaultInputOptions().apply(opts...)
	out := make(chan string, messageBoxBufferSize)
	modal := newMessageBoxModal(options, out, p.styles)
	if err := p.send(msgModal{modal: modal}); err != nil {
		return nil, err
	}
	return &MessageBoxHandle{
		prompter: p,
		modal:    modal,
		stream:   out,
	}, nil
}

func (h *MessageBoxHandle) Messages() <-chan string {
	return h.stream
}

func (h *MessageBoxHandle) Clear() error {
	return h.prompter.send(msgMessageBoxClear{modal: h.modal})
}
