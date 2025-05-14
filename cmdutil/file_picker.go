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

type file_picker_model struct {
	// The fp model
	filepicker filepicker.Model

	// List of files that the user selected
	SelectedFiles []string

	// So we can have a title with the fp
	// This is so we can explain what purpose the files are being selected for
	title string

	// To store the quit signal
	quitting bool

	// To handle possible errors
	err error
}

type clearErrorMsg struct{}

// We want to show a temporary error if a file is selected more than once
func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m file_picker_model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m file_picker_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		// If user selects a path already in selected files, display an error
		if slices.Contains(m.SelectedFiles, path) {
			m.err = errors.New(
				"you have this file already. kindly select another file or save your selection",
			)
			return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
		} else {
			m.SelectedFiles = append(m.SelectedFiles, path)
		}
	}

	return m, cmd
}

func (m file_picker_model) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder

	s.WriteString("\n  ")
	switch {
	case m.err != nil:
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	case len(m.SelectedFiles) == 0:
		s.WriteString(m.title)
	default:
		s.WriteString(
			"Selected files: " + m.filepicker.Styles.Selected.Render(
				strconv.Itoa(len(m.SelectedFiles))),
		)
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	s.WriteString("You can use ctrl+c, q or s to quit and save your selection.")

	return s.String()
}

func initialModel(dir string, title string, allowed_types []string) file_picker_model {
	fp := filepicker.New()
	fp.CurrentDirectory = dir
	fp.AllowedTypes = allowed_types

	im := file_picker_model{
		filepicker: fp,
		title:      title,
	}

	return im
}

func FileTreeSelect(dir string, title string, allowed_types []string) file_picker_model {
	// We want the option of excluding already select files.
	// We can pass exact filepaths to allowed_types to ensure only these are
	// selectable.
	m := initialModel(dir, title, allowed_types)

	tm, _ := tea.NewProgram(&m).Run()
	mm := tm.(file_picker_model)
	fmt.Println("\n  You selected: " + m.filepicker.Styles.Selected.Render(
		strconv.Itoa(len(mm.SelectedFiles))+" file(s) in total.") + "\n",
	)

	// Return file_picker_model struct after user exits or saves.
	return mm
}
