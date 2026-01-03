package prompter

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/olimci/prompter/promise"
)

func newConfirmModal(prompt string, confirm bool, resolve promise.Resolver[bool], styles *Styles) modal {
	return &confirmModal{
		prompt:  prompt,
		confirm: confirm,
		resolve: resolve,
		styles:  styles,
	}
}

type confirmModal struct {
	prompt  string
	confirm bool
	resolve promise.Resolver[bool]
	styles  *Styles
}

func (m *confirmModal) Init() tea.Cmd {
	return nil
}

func (m *confirmModal) Update(msg tea.Msg) (modal, tea.Cmd, string) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.confirm = true
			m.resolve.Success(true)
			return nil, nil, m.final()
		case "n", "N":
			m.confirm = false
			m.resolve.Success(false)
			return nil, nil, m.final()
		case "enter":
			m.resolve.Success(m.confirm)
			return nil, nil, m.final()
		case "left", "right", "up", "down", "tab", "space":
			m.confirm = !m.confirm
			return m, nil, ""

		}
	}
	return m, nil, ""
}

func (m *confirmModal) View() string {
	prompt := m.styles.Confirm.Prompt.Render(m.prompt)

	yes := "y"
	no := "n"
	if m.confirm {
		yes = "Y"
	} else {
		no = "N"
	}
	if m.confirm {
		yes = m.styles.Confirm.ActiveOption.Render(yes)
		no = m.styles.Confirm.InactiveOption.Render(no)
	} else {
		yes = m.styles.Confirm.InactiveOption.Render(yes)
		no = m.styles.Confirm.ActiveOption.Render(no)
	}

	selected := fmt.Sprintf("[%s/%s]", yes, no)
	return fmt.Sprintf("%s %s", prompt, selected)
}

func (m *confirmModal) final() string {
	prompt := m.styles.Confirm.Prompt.Render(m.prompt)

	selected := "NO"
	if m.confirm {
		selected = "YES"
	}
	selected = m.styles.Confirm.Final.Render(selected)

	return fmt.Sprintf("%s %s", prompt, selected)
}
