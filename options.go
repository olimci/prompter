package prompter

import "context"

func defaultOptions() *options {
	return &options{
		context: context.Background(),
		styles:  DefaultStyles(),
	}
}

type options struct {
	context       context.Context
	styles        *Styles
	maxHistoryLen int
	viewFunc      viewFunc
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}

	return o
}

type Option func(*options)

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.context = ctx
	}
}

func WithStyles(styles *Styles) Option {
	return func(o *options) {
		if styles == nil {
			panic("styles cannot be nil")
		}
		o.styles = styles
	}
}

func WithMaxHistoryLen(max int) Option {
	return func(o *options) {
		if max < 0 {
			panic("max history length cannot be negative")
		}
		o.maxHistoryLen = max
	}
}

func WithViewFunc(view func(history []string, modal string) string) Option {
	return func(o *options) {
		o.viewFunc = view
	}
}
