package prompter

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type msgMessageBoxClear struct {
	modal *messageBoxModal
}

func newMessageBoxModal(opts *inputOptions, out chan string, styles *Styles) *messageBoxModal {
	in := textinput.New()
	if opts.promptSet {
		in.Prompt = opts.prompt
	}
	in.Validate = opts.validate
	in.CharLimit = opts.charLimit
	in.Placeholder = opts.placeholder
	if opts.width > 0 {
		in.Width = opts.width
	}
	if len(opts.suggestions) > 0 {
		in.SetSuggestions(opts.suggestions)
	}
	in.ShowSuggestions = len(opts.suggestions) > 0
	in.Focus()
	in.PromptStyle = styles.Input.Prompt
	in.TextStyle = styles.Input.Text
	in.PlaceholderStyle = styles.Input.Placeholder
	in.Cursor.Style = styles.Input.Cursor

	return &messageBoxModal{
		in:     in,
		out:    out,
		opts:   opts,
		styles: styles,
	}
}

type messageBoxModal struct {
	in     textinput.Model
	out    chan string
	closed bool
	opts   *inputOptions
	styles *Styles
}

func (m *messageBoxModal) Init() tea.Cmd {
	return nil
}

func (m *messageBoxModal) Update(msg tea.Msg) (modal, tea.Cmd, string) {
	switch msg := msg.(type) {
	case msgMessageBoxClear:
		if msg.modal != m {
			return m, nil, ""
		}
		m.close()
		return nil, nil, ""
	case tea.KeyMsg:
		if msg.String() == "enter" {
			if m.in.Err != nil {
				return m, nil, ""
			}
			value := m.in.Value()
			if value != "" {
				m.send(value)
			}
			m.in.SetValue("")
			return m, nil, ""
		}
	}

	var cmd tea.Cmd
	m.in, cmd = m.in.Update(msg)
	return m, cmd, ""
}

func (m *messageBoxModal) View() string {
	if m.in.Err != nil {
		errMsg := m.styles.Input.Error.Render(m.in.Err.Error())
		return fmt.Sprintf("%s %s", m.in.View(), errMsg)
	}
	return m.in.View()
}

func (m *messageBoxModal) send(value string) {
	select {
	case m.out <- value:
	default:
	}
}

func (m *messageBoxModal) close() {
	if m.closed {
		return
	}
	m.closed = true
	close(m.out)
}
