package ui

import "github.com/charmbracelet/lipgloss"

var (
	TextBar    = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).MaxHeight(5).Render
	TimerStyle = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).MarginTop(2).Render
	WpmScore   = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render

	Correct    = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Render
	Wrong      = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render
	SpaceWrong = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render("-")
	UnTyped    = lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Render
	Cursor     = lipgloss.NewStyle().Foreground(lipgloss.Color("243")).Underline(true).Render
)
