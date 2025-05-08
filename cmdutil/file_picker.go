package cmdutil

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	filepicker    filepicker.Model
	SelectedFiles []string
	quitting      bool
	err           error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "s":
			m.quitting = true
			return m, tea.Quit
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		// If user selects a path already in selected files, display an
		// error to the user and clear.
		// Get the path of the selected file.
		if slices.Contains(m.SelectedFiles, path) {
			m.err = errors.New("you have this file already. kindly select another file or save your selection")
			// Can we do this without returning?
			return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
		} else {
			m.SelectedFiles = append(m.SelectedFiles, path)
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if len(m.SelectedFiles) == 0 {
		s.WriteString("Pick a file:")
	} else {
		s.WriteString(
			"Selected files: " + m.filepicker.Styles.Selected.Render(
				strconv.Itoa(len(m.SelectedFiles))),
		)
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	s.WriteString("You can use ctrl+c, q or s to quit and save your selection.")
	return s.String()
}

func FileTreeSelect(dir string) model {
	fp := filepicker.New()
	fp.CurrentDirectory = dir

	m := model{
		filepicker: fp,
	}
	tm, _ := tea.NewProgram(&m).Run()
	mm := tm.(model)
	fmt.Println("\n  You selected: " + m.filepicker.Styles.Selected.Render(
		strconv.Itoa(len(mm.SelectedFiles))+" file(s) in total.") + "\n",
	)

	// Return model struct after user exits or saves.
	return mm
}
