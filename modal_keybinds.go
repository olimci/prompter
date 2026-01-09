package prompter

import (
	"fmt"
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

type msgKeybindClear struct {
	modal *keybindModal
}

type msgKeybindSet struct {
	modal    *keybindModal
	bindings []Keybind
}

func newKeybindModal(opts *keybindOptions, bindings []Keybind, out chan KeybindEvent, styles *Styles) *keybindModal {
	m := &keybindModal{
		out:    out,
		opts:   opts,
		styles: styles,
	}
	m.setBindings(bindings)
	return m
}

type keybindModal struct {
	bindings      []Keybind
	bindingsByKey map[string]Keybind
	out           chan KeybindEvent
	closed        bool
	opts          *keybindOptions
	styles        *Styles
}

func (m *keybindModal) Init() tea.Cmd {
	return nil
}

func (m *keybindModal) Update(msg tea.Msg) (modal, tea.Cmd, string) {
	switch msg := msg.(type) {
	case msgKeybindClear:
		if msg.modal != m {
			return m, nil, ""
		}
		m.close()
		return nil, nil, ""

	case msgKeybindSet:
		if msg.modal != m {
			return m, nil, ""
		}
		m.setBindings(msg.bindings)
		return m, nil, ""

	case tea.KeyMsg:
		if binding, ok := m.match(msg.String()); ok {
			m.send(KeybindEvent{Key: binding.Key, Event: binding.Event})
			return m, nil, ""
		}
	}

	return m, nil, ""
}

func (m *keybindModal) View() string {
	if len(m.bindings) == 0 {
		return ""
	}

	parts := make([]string, 0, len(m.bindings))
	for _, b := range m.bindings {
		label := b.Description
		if label == "" {
			label = b.Event
		}
		key := m.styles.Select.Cursor.Render(fmt.Sprintf("[%s]", b.Key))
		desc := m.styles.Select.Option.Render(label)
		parts = append(parts, fmt.Sprintf("%s %s", key, desc))
	}

	line := strings.Join(parts, m.opts.separator)
	if m.opts.promptSet && m.opts.prompt != "" {
		prompt := m.styles.Select.Prompt.Render(m.opts.prompt)
		return fmt.Sprintf("%s %s", prompt, line)
	}
	return line
}

func (m *keybindModal) send(ev KeybindEvent) {
	select {
	case m.out <- ev:
	default:
	}
}

func (m *keybindModal) close() {
	if m.closed {
		return
	}
	m.closed = true
	close(m.out)
}

func (m *keybindModal) setBindings(bindings []Keybind) {
	m.bindings = make([]Keybind, 0, len(bindings))
	m.bindingsByKey = make(map[string]Keybind, len(bindings))
	for _, b := range bindings {
		if b.Key == "" || b.Event == "" {
			continue
		}
		normalizedKey := normalizeKey(b.Key)
		b.Key = normalizedKey
		m.bindings = append(m.bindings, b)
		m.bindingsByKey[normalizedKey] = b
	}
}

func (m *keybindModal) match(key string) (Keybind, bool) {
	if m.bindingsByKey == nil {
		return Keybind{}, false
	}
	if binding, ok := m.bindingsByKey[normalizeKey(key)]; ok {
		return binding, true
	}
	return Keybind{}, false
}

func normalizeKey(key string) string {
	key = strings.TrimSpace(key)
	if key == "" {
		return ""
	}
	runes := []rune(key)
	if len(runes) == 1 && unicode.IsLetter(runes[0]) {
		return strings.ToLower(key)
	}
	return key
}
