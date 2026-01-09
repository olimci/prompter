package prompter

import (
	tea "github.com/charmbracelet/bubbletea"
)

type msgStatusKeybindClear struct {
	modal *statusKeybindModal
}

func newStatusKeybindModal(status *statusModal, keybinds *keybindModal) *statusKeybindModal {
	return &statusKeybindModal{
		status:   status,
		keybinds: keybinds,
	}
}

type statusKeybindModal struct {
	status   *statusModal
	keybinds *keybindModal
}

func (m *statusKeybindModal) Init() tea.Cmd {
	if m.status == nil {
		return nil
	}
	return m.status.Init()
}

func (m *statusKeybindModal) Update(msg tea.Msg) (modal, tea.Cmd, string) {
	switch msg := msg.(type) {
	case msgStatusKeybindClear:
		if msg.modal != m {
			return m, nil, ""
		}
		if m.keybinds != nil {
			m.keybinds.close()
		}
		if m.status == nil {
			return nil, nil, ""
		}
		return nil, nil, m.status.final()
	case msgStatusClear:
		if m.status != nil && msg.modal == m.status {
			if m.keybinds != nil {
				m.keybinds.close()
			}
			return nil, nil, m.status.final()
		}
	case msgKeybindClear:
		if m.keybinds != nil && msg.modal == m.keybinds {
			m.keybinds.close()
			if m.status == nil {
				return nil, nil, ""
			}
			return nil, nil, m.status.final()
		}
	}

	var statusCmd tea.Cmd
	if m.status != nil {
		next, cmd, final := m.status.Update(msg)
		if next == nil {
			if m.keybinds != nil {
				m.keybinds.close()
			}
			return nil, nil, final
		}
		if updated, ok := next.(*statusModal); ok {
			m.status = updated
		}
		statusCmd = cmd
	}

	var keybindCmd tea.Cmd
	if m.keybinds != nil {
		next, cmd, final := m.keybinds.Update(msg)
		if next == nil {
			m.keybinds.close()
			return nil, nil, final
		}
		if updated, ok := next.(*keybindModal); ok {
			m.keybinds = updated
		}
		keybindCmd = cmd
	}

	return m, tea.Batch(statusCmd, keybindCmd), ""
}

func (m *statusKeybindModal) View() string {
	statusView := ""
	if m.status != nil {
		statusView = m.status.View()
	}
	keybindView := ""
	if m.keybinds != nil {
		keybindView = m.keybinds.View()
	}

	switch {
	case statusView == "" && keybindView == "":
		return ""
	case statusView == "":
		return keybindView
	case keybindView == "":
		return statusView
	default:
		return statusView + "\n" + keybindView
	}
}
