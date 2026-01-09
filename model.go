package prompter

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type msgLog struct {
	message string
}

type msgModal struct {
	modal modal
}

type msgClearHistory struct{}

func newModel(o *options) *model {
	history := make([]string, 0)
	if o.maxHistoryLen > 0 {
		history = make([]string, 0, o.maxHistoryLen)
	}
	return &model{
		history:       history,
		modalQueue:    make([]msgModal, 0),
		maxHistoryLen: o.maxHistoryLen,
		viewFunc:      o.viewFunc,
	}
}

type viewFunc func(history []string, modal string) string

type model struct {
	history       []string
	modal         modal
	modalQueue    []msgModal
	ready         chan<- struct{}
	maxHistoryLen int
	viewFunc      viewFunc
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case msgLog:
		m.appendHistory(msg.message)
		return m, nil

	case msgModal:
		if m.modal == nil {
			m.modal = msg.modal
			return m, m.modal.Init()
		} else {
			m.modalQueue = append(m.modalQueue, msg)
		}

	case msgClearHistory:
		m.clearHistory()
		return m, nil
	}

	if m.modal != nil {
		next, cmd, final := m.modal.Update(msg)
		if m.modal = next; next == nil {
			if final != "" {
				m.appendHistory(final)
			}
			if m.modal = m.popModal(); m.modal != nil {
				return m, tea.Batch(cmd, m.modal.Init())
			}
		}
		return m, cmd
	}

	return m, nil
}

func (m *model) View() string {
	if m.viewFunc != nil {
		modalView := ""
		if m.modal != nil {
			modalView = m.modal.View()
		}
		return m.viewFunc(m.history, modalView)
	}

	lines := make([]string, 0, len(m.history)+1)
	for _, line := range m.history {
		lines = append(lines, line)
	}
	if m.modal != nil {
		lines = append(lines, m.modal.View())
	}
	return strings.Join(lines, "\n") + "\n"
}

func (m *model) popModal() modal {
	if len(m.modalQueue) == 0 {
		return nil
	}
	modal := m.modalQueue[0]
	m.modalQueue = m.modalQueue[1:]
	return modal.modal
}

func (m *model) appendHistory(line string) {
	m.history = append(m.history, line)
	if m.maxHistoryLen > 0 && len(m.history) > m.maxHistoryLen {
		m.history = m.history[len(m.history)-m.maxHistoryLen:]
	}
}

func (m *model) clearHistory() {
	m.history = m.history[:0]
}
