package prompter

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Confirm ConfirmStyles
	Input   InputStyles
	Select  SelectStyles
	Status  StatusStyles
}

type ConfirmStyles struct {
	Prompt         lipgloss.Style
	ActiveOption   lipgloss.Style
	InactiveOption lipgloss.Style
	Final          lipgloss.Style
}

type InputStyles struct {
	Prompt      lipgloss.Style
	Text        lipgloss.Style
	Placeholder lipgloss.Style
	Cursor      lipgloss.Style
	Error       lipgloss.Style
	Final       lipgloss.Style
}

type SelectStyles struct {
	Prompt         lipgloss.Style
	Cursor         lipgloss.Style
	Option         lipgloss.Style
	SelectedOption lipgloss.Style
	Final          lipgloss.Style
}

type StatusStyles struct {
	IdleIcon     lipgloss.Style
	Spinner      lipgloss.Style
	Message      lipgloss.Style
	FinalIcon    lipgloss.Style
	FinalMessage lipgloss.Style
}

func DefaultStyles() *Styles {
	return &Styles{
		Confirm: ConfirmStyles{
			Prompt:         lipgloss.NewStyle().Bold(true),
			ActiveOption:   lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true),
			InactiveOption: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
			Final:          lipgloss.NewStyle().Foreground(lipgloss.Color("244")),
		},
		Input: InputStyles{
			Prompt:      lipgloss.NewStyle().Bold(true),
			Text:        lipgloss.NewStyle(),
			Placeholder: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
			Cursor:      lipgloss.NewStyle().Foreground(lipgloss.Color("205")),
			Error:       lipgloss.NewStyle().Foreground(lipgloss.Color("196")),
			Final:       lipgloss.NewStyle().Foreground(lipgloss.Color("244")),
		},
		Select: SelectStyles{
			Prompt:         lipgloss.NewStyle().Bold(true),
			Cursor:         lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true),
			Option:         lipgloss.NewStyle(),
			SelectedOption: lipgloss.NewStyle().Bold(true),
			Final:          lipgloss.NewStyle().Foreground(lipgloss.Color("244")),
		},
		Status: StatusStyles{
			IdleIcon:     lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
			Spinner:      lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true),
			Message:      lipgloss.NewStyle(),
			FinalIcon:    lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true),
			FinalMessage: lipgloss.NewStyle().Foreground(lipgloss.Color("244")),
		},
	}
}
