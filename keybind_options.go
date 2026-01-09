package prompter

type keybindOptions struct {
	prompt    string
	promptSet bool
	separator string
}

type KeybindOption func(*keybindOptions)

func defaultKeybindOptions() *keybindOptions {
	return &keybindOptions{
		separator: " • ",
	}
}

func (o *keybindOptions) apply(opts ...KeybindOption) *keybindOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithKeybindPrompt(prompt string) KeybindOption {
	return func(o *keybindOptions) {
		o.prompt = prompt
		o.promptSet = true
	}
}

func WithKeybindSeparator(separator string) KeybindOption {
	return func(o *keybindOptions) {
		o.separator = separator
	}
}
