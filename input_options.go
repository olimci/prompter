package prompter

type inputOptions struct {
	validate    func(string) error
	charLimit   int
	placeholder string
	prompt      string
	promptSet   bool
	suggestions []string
	width       int
}

type InputOption func(*inputOptions)

func defaultInputOptions() *inputOptions {
	return &inputOptions{}
}

func (o *inputOptions) apply(opts ...InputOption) *inputOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithInputValidate(validate func(string) error) InputOption {
	return func(o *inputOptions) {
		o.validate = validate
	}
}

func WithInputCharLimit(limit int) InputOption {
	return func(o *inputOptions) {
		o.charLimit = limit
	}
}

func WithInputPlaceholder(placeholder string) InputOption {
	return func(o *inputOptions) {
		o.placeholder = placeholder
	}
}

func WithInputPrompt(prompt string) InputOption {
	return func(o *inputOptions) {
		o.prompt = prompt
		o.promptSet = true
	}
}

func WithInputSuggestions(suggestions []string) InputOption {
	return func(o *inputOptions) {
		o.suggestions = suggestions
	}
}

func WithInputWidth(width int) InputOption {
	return func(o *inputOptions) {
		o.width = width
	}
}
