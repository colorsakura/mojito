package tui

import "github.com/charmbracelet/lipgloss"

// TODO: implement tui theme function
var (
	appStyle = lipgloss.NewStyle().
		Width(60).
		Height(15).
		Foreground(lipgloss.Color("#31363b")).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#3daee9"))

	//itemHight = lipgloss.NewStyle().Bold(true)
	//playerStyle = lipgloass.NewStyle()
)
