package prompter

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/olimci/prompter/promise"
)

func newInputModal(opts *inputOptions, resolve promise.Resolver[string], styles *Styles) modal {
	return &inputModal{
		prompt:  opts.prompt,
		in:      textinput.New(),
		options: opts,
		resolve: resolve,
		styles:  styles,
	}
}

type inputModal struct {
	prompt  string
	in      textinput.Model
	options *inputOptions
	resolve promise.Resolver[string]
	styles  *Styles
}

func (m *inputModal) Init() tea.Cmd {
	if m.options.promptSet {
		m.in.Prompt = m.prompt
	}
	m.in.Validate = m.options.validate
	m.in.CharLimit = m.options.charLimit
	m.in.Placeholder = m.options.placeholder
	if m.options.width > 0 {
		m.in.Width = m.options.width
	}
	if len(m.options.suggestions) > 0 {
		m.in.SetSuggestions(m.options.suggestions)
	}
	m.in.ShowSuggestions = len(m.options.suggestions) > 0
	m.in.Focus()
	m.in.PromptStyle = m.styles.Input.Prompt
	m.in.TextStyle = m.styles.Input.Text
	m.in.PlaceholderStyle = m.styles.Input.Placeholder
	m.in.Cursor.Style = m.styles.Input.Cursor
	return nil
}

func (m *inputModal) Update(msg tea.Msg) (modal, tea.Cmd, string) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.in.Err != nil {
				return m, nil, ""
			}
			m.resolve.Success(m.in.Value())
			return nil, nil, m.final()
		}
	}

	var cmd tea.Cmd
	m.in, cmd = m.in.Update(msg)
	return m, cmd, ""
}

func (m *inputModal) View() string {
	if m.in.Err != nil {
		errMsg := m.styles.Input.Error.Render(m.in.Err.Error())
		return fmt.Sprintf("%s %s", m.in.View(), errMsg)
	}
	return fmt.Sprintf("%s", m.in.View())
}

func (m *inputModal) final() string {
	prompt := m.styles.Input.Prompt.Render(m.prompt)
	value := m.styles.Input.Final.Render(m.in.Value())
	return fmt.Sprintf("%s%s", prompt, value)
}
