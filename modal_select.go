package prompter

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/olimci/prompter/promise"
)

func newSelectModal(prompt string, options []string, index int, resolve promise.Resolver[string], styles *Styles) modal {
	return &selectModal{
		prompt:  prompt,
		options: options,
		index:   index,
		resolve: resolve,
		styles:  styles,
	}
}

type selectModal struct {
	prompt  string
	options []string
	index   int
	resolve promise.Resolver[string]
	styles  *Styles
}

func (m *selectModal) Init() tea.Cmd {
	return nil
}

func (m *selectModal) Update(msg tea.Msg) (modal, tea.Cmd, string) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if len(m.options) == 0 {
				return m, nil, ""
			}
			m.index--
			if m.index < 0 {
				m.index = len(m.options) - 1
			}
			return m, nil, ""
		case "down", "j":
			if len(m.options) == 0 {
				return m, nil, ""
			}
			m.index++
			if m.index >= len(m.options) {
				m.index = 0
			}
			return m, nil, ""
		case "enter":
			if len(m.options) == 0 {
				return m, nil, ""
			}
			selected := m.options[m.index]
			m.resolve.Success(selected)
			return nil, nil, m.final(selected)
		}
	}

	return m, nil, ""
}

func (m *selectModal) View() string {
	var b strings.Builder
	if m.prompt != "" {
		prompt := m.styles.Select.Prompt.Render(m.prompt)
		b.WriteString(prompt)
		b.WriteString("\n")
	}

	for i, opt := range m.options {
		cursor := "  "
		if i == m.index {
			cursor = "> "
		}
		option := opt
		if i == m.index {
			cursor = m.styles.Select.Cursor.Render(cursor)
			option = m.styles.Select.SelectedOption.Render(option)
		} else {
			option = m.styles.Select.Option.Render(option)
		}
		b.WriteString(cursor)
		b.WriteString(option)
		if i < len(m.options)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func (m *selectModal) final(selected string) string {
	if m.prompt == "" {
		return m.styles.Select.Final.Render(selected)
	}
	prompt := m.styles.Select.Prompt.Render(m.prompt)
	selected = m.styles.Select.Final.Render(selected)
	return fmt.Sprintf("%s %s", prompt, selected)
}
