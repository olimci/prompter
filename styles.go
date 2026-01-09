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
	ErrorIcon    lipgloss.Style
	ErrorMessage lipgloss.Style
}

func DefaultStyles() *Styles {
	text := lipgloss.AdaptiveColor{Light: "#24292F", Dark: "#C0CAF5"}
	muted := lipgloss.AdaptiveColor{Light: "#6E7781", Dark: "#565F89"}
	faint := lipgloss.AdaptiveColor{Light: "#8C959F", Dark: "#3B4261"}
	accent := lipgloss.AdaptiveColor{Light: "#0969DA", Dark: "#7AA2F7"}
	success := lipgloss.AdaptiveColor{Light: "#1A7F37", Dark: "#9ECE6A"}
	errorCol := lipgloss.AdaptiveColor{Light: "#CF222E", Dark: "#F7768E"}

	return &Styles{
		Confirm: ConfirmStyles{
			Prompt:         lipgloss.NewStyle().Foreground(accent).Bold(true),
			ActiveOption:   lipgloss.NewStyle().Foreground(success).Bold(true),
			InactiveOption: lipgloss.NewStyle().Foreground(muted),
			Final:          lipgloss.NewStyle().Foreground(muted),
		},
		Input: InputStyles{
			Prompt:      lipgloss.NewStyle().Foreground(accent).Bold(true),
			Text:        lipgloss.NewStyle().Foreground(text),
			Placeholder: lipgloss.NewStyle().Foreground(faint),
			Cursor:      lipgloss.NewStyle().Foreground(accent),
			Error:       lipgloss.NewStyle().Foreground(errorCol).Bold(true),
			Final:       lipgloss.NewStyle().Foreground(muted),
		},
		Select: SelectStyles{
			Prompt:         lipgloss.NewStyle().Foreground(accent).Bold(true),
			Cursor:         lipgloss.NewStyle().Foreground(accent).Bold(true),
			Option:         lipgloss.NewStyle().Foreground(text),
			SelectedOption: lipgloss.NewStyle().Foreground(accent).Bold(true),
			Final:          lipgloss.NewStyle().Foreground(muted),
		},
		Status: StatusStyles{
			IdleIcon:     lipgloss.NewStyle().Foreground(muted),
			Spinner:      lipgloss.NewStyle().Foreground(accent).Bold(true),
			Message:      lipgloss.NewStyle().Foreground(text),
			FinalIcon:    lipgloss.NewStyle().Foreground(success).Bold(true),
			FinalMessage: lipgloss.NewStyle().Foreground(muted),
			ErrorIcon:    lipgloss.NewStyle().Foreground(errorCol).Bold(true),
			ErrorMessage: lipgloss.NewStyle().Foreground(errorCol),
		},
	}
}
