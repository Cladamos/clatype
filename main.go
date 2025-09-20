package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"

	"github.com/cladamos/clatype/components"
	"github.com/cladamos/clatype/ui"
)

type Model struct {
	width, height int
	timer         timer.Model
	wordData      string
	keymap        keymap
	help          help.Model
	textInput     textinput.Model
}

type keymap struct {
	restart key.Binding
	quit    key.Binding
}

var timeout = time.Second * 30

func main() {

	m := initialModel()
	m.keymap.restart.SetEnabled(false)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() Model {
	ti := textinput.New()
	ti.Width = 0
	ti.Focus()

	return Model{
		timer: timer.NewWithInterval(timeout, time.Second),
		keymap: keymap{
			quit:    key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
			restart: key.NewBinding(key.WithKeys("ctrl+r"), key.WithHelp("ctrl+r", "restart")),
		},
		help:      help.New(),
		wordData:  components.GenerateWords(),
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Cmd(
		m.timer.Init(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case string:
		m.wordData = msg

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit

		case key.Matches(msg, m.keymap.restart):
			newM := initialModel()
			newM.width = m.width
			newM.height = m.height
			return newM, newM.timer.Init()
		}
	case timer.TickMsg:

		m.timer, cmd = m.timer.Update(msg)

		if m.timer.Timedout() {
			m.keymap.restart.SetEnabled(true)
			m.textInput.Blur()
		}
		return m, cmd
	}
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m Model) helpView() string {
	return lipgloss.NewStyle().Align(lipgloss.Center).MarginTop(2).Render(
		m.help.ShortHelpView([]key.Binding{
			m.keymap.restart,
			m.keymap.quit,
		}),
	)
}

func (m Model) View() string {

	if len(m.wordData) == 0 {
		return "Loading words..."
	}

	s := lipgloss.JoinVertical(lipgloss.Center, m.renderInput(), ui.TimerStyle(m.timer.View()))

	if m.timer.Timedout() {
		wpm, accuracy := components.CalculateWpm(m.wordData, m.textInput.Value(), timeout)

		s = lipgloss.JoinVertical(lipgloss.Center,
			ui.WpmScore("Your wpm is "+strconv.Itoa(int(wpm))),
			ui.WpmScore("Your accuracy is "+strconv.Itoa(int(accuracy))+"%"),
			m.helpView(),
		)
	}
	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		s,
	)
}

func (m Model) renderInput() string {
	userInput := m.textInput.Value()
	target := m.wordData

	styled := ""
	for i, r := range userInput {
		if i < len(target) && r == rune(target[i]) {
			styled += ui.Correct(string(r))
		} else {
			if unicode.IsSpace(rune(target[i])) {
				styled += ui.SpaceWrong
			} else {
				styled += ui.Wrong(string(target[i]))
			}
		}

	}
	if len(userInput) < len(target) {
		remaining := target[len(userInput):]
		styled += ui.Cursor(string(remaining[0]))
		styled += ui.UnTyped(remaining[1:])
	}

	wrapped := wordwrap.String(styled, 50)
	lines := strings.Split(wrapped, "\n")

	cursorLine := len(userInput) / 50
	maxLines := 3

	start := cursorLine - 1
	if start < 0 {
		start = 0
	}
	end := start + maxLines

	visibleLines := lines[start:end]

	styledLines := []string{}
	for _, l := range visibleLines {
		styledLines = append(styledLines, ui.UnTyped(l))
	}

	return lipgloss.JoinVertical(lipgloss.Center, styledLines...)

}
