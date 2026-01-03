package prompter

import tea "github.com/charmbracelet/bubbletea"

type modal interface {
	// Init initialises the modal and returns a command to be executed
	Init() tea.Cmd
	// Update updates the modal,
	// - nil next return means that the modal is done,
	// - final represents the final view for the modal
	Update(tea.Msg) (next modal, cmd tea.Cmd, final string)
	// View returns the modals view
	View() string
}
