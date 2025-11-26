package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://duckduckgo.com"

type (
	statusMsg int
	errMsg    struct{ err error }
)

type model struct {
	status uint8
	err    error
}

func main() {
	if err := tea.NewProgram(model{}).Start(); err != nil {
		log.Fatalf("Uh oh, there was an error: %v", err)
	}
}

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.status = uint8(msg)
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
	}

	s := fmt.Sprintf("Checking %s...", url)

	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(int(m.status)))
	}

	return "\n" + s + "\n\n"
}

func checkServer() tea.Msg {
	c := &http.Client{Timeout: time.Second * 10}

	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}

	return statusMsg(res.StatusCode)
}

func (e errMsg) Error() string { return e.err.Error() }
