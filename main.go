package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Style struct {
	containerStyle lipgloss.Style
	inputStyle     lipgloss.Style
	phraseStyle    lipgloss.Style
	errorStyle     lipgloss.Style
	cursorStyle    lipgloss.Style
}

func main() {
	p := tea.NewProgram(NewTypingModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
