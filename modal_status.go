package prompter

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type statusState int

const (
	statusIdle statusState = iota
	statusWorking
	statusProgress
	statusSuccess
	statusError
)

const (
	statusIdleIcon    = "○"
	statusSuccessIcon = "✓"
	statusErrorIcon   = "✗"
)

type msgStatusIdle struct {
	modal   *statusModal
	message string
}

type msgStatusWorking struct {
	modal   *statusModal
	message string
}

type msgStatusProgress struct {
	modal   *statusModal
	message string
	percent float64
}

type msgStatusProgressValue struct {
	modal   *statusModal
	percent float64
}

type msgStatusMessage struct {
	modal   *statusModal
	message string
}

type msgStatusSuccess struct {
	modal   *statusModal
	message string
}

type msgStatusError struct {
	modal   *statusModal
	message string
}

type msgStatusClear struct {
	modal *statusModal
}

func newStatusModal(state statusState, message string, styles *Styles) *statusModal {
	spin := spinner.New()
	spin.Spinner = spinner.Pulse
	spin.Style = styles.Status.Spinner

	bar := progress.New(
		progress.WithWidth(32),
		progress.WithDefaultGradient(),
		progress.WithFillCharacters('█', '░'),
	)
	bar.PercentageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	return &statusModal{
		state:       state,
		message:     message,
		spinner:     spin,
		progressBar: bar,
		styles:      styles,
	}
}

type statusModal struct {
	state       statusState
	message     string
	progress    float64
	spinner     spinner.Model
	progressBar progress.Model
	styles      *Styles
}

func (m *statusModal) Init() tea.Cmd {
	switch m.state {
	case statusWorking:
		return m.spinner.Tick
	case statusProgress:
		return m.progressBar.SetPercent(m.progress)
	default:
		return nil
	}
}

func (m *statusModal) Update(msg tea.Msg) (modal, tea.Cmd, string) {
	switch msg := msg.(type) {
	case msgStatusIdle:
		if msg.modal != m {
			return m, nil, ""
		}
		m.state = statusIdle
		m.message = msg.message
		return m, nil, ""
	case msgStatusWorking:
		if msg.modal != m {
			return m, nil, ""
		}
		m.state = statusWorking
		m.message = msg.message
		return m, m.spinner.Tick, ""
	case msgStatusProgress:
		if msg.modal != m {
			return m, nil, ""
		}
		m.state = statusProgress
		m.message = msg.message
		m.progress = msg.percent
		return m, m.progressBar.SetPercent(m.progress), ""
	case msgStatusProgressValue:
		if msg.modal != m {
			return m, nil, ""
		}
		m.state = statusProgress
		m.progress = msg.percent
		return m, m.progressBar.SetPercent(m.progress), ""
	case msgStatusMessage:
		if msg.modal != m {
			return m, nil, ""
		}
		m.message = msg.message
		return m, nil, ""
	case msgStatusSuccess:
		if msg.modal != m {
			return m, nil, ""
		}
		m.state = statusSuccess
		m.message = msg.message
		return m, nil, ""
	case msgStatusError:
		if msg.modal != m {
			return m, nil, ""
		}
		m.state = statusError
		m.message = msg.message
		return m, nil, ""
	case msgStatusClear:
		if msg.modal != m {
			return m, nil, ""
		}
		return nil, nil, m.final()
	}

	switch m.state {
	case statusWorking:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd, ""
	case statusProgress:
		var cmd tea.Cmd
		var model tea.Model
		model, cmd = m.progressBar.Update(msg)
		if progressModel, ok := model.(progress.Model); ok {
			m.progressBar = progressModel
		}
		return m, cmd, ""
	default:
		return m, nil, ""
	}
}

func (m *statusModal) View() string {
	indicator := m.indicator()
	message := m.messageView()
	if m.message == "" {
		return indicator
	}
	return fmt.Sprintf("%s %s", indicator, message)
}

func (m *statusModal) indicator() string {
	switch m.state {
	case statusWorking:
		return m.spinner.View()
	case statusProgress:
		return m.progressBar.View()
	case statusSuccess:
		return m.styles.Status.FinalIcon.Render(statusSuccessIcon)
	case statusError:
		return m.styles.Status.ErrorIcon.Render(statusErrorIcon)
	default:
		return m.styles.Status.IdleIcon.Render(statusIdleIcon)
	}
}

func (m *statusModal) final() string {
	switch m.state {
	case statusSuccess, statusError:
		return m.View()
	default:
		return ""
	}
}

func (m *statusModal) messageView() string {
	switch m.state {
	case statusSuccess:
		return m.styles.Status.FinalMessage.Render(m.message)
	case statusError:
		return m.styles.Status.ErrorMessage.Render(m.message)
	default:
		return m.styles.Status.Message.Render(m.message)
	}
}
