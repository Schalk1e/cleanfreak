package cmdutil

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type result_model struct {
	desc    string
	choices []string
	cursor  int
	Choice  string
}

func (m result_model) Init() tea.Cmd {
	return nil
}

func (m result_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.Choice = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m result_model) View() string {
	s := strings.Builder{}
	s.WriteString(m.desc)
	s.WriteString("\n")

	for i, choice := range m.choices {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choice)
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func ListResult(choices []string, desc string) result_model {
	p := tea.NewProgram(
		result_model{
			choices: choices,
			desc:    desc,
		},
	)

	tm, _ := p.Run()
	mm := tm.(result_model)

	return mm
}
