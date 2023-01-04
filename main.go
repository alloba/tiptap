package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ViewStyle struct {
	containerStyle lipgloss.Style
	inputStyle     lipgloss.Style
	phraseStyle    lipgloss.Style
	errorStyle     lipgloss.Style
	cursorStyle    lipgloss.Style
}

func main() {

	//TODO: this should be loaded from a config file eventually.
	var globalStyle = ViewStyle{
		containerStyle: lipgloss.NewStyle().Bold(true).PaddingTop(1).PaddingLeft(2),
		inputStyle:     lipgloss.NewStyle().Background(lipgloss.Color("#16001E")).Foreground(lipgloss.Color("#DE639A")),
		phraseStyle:    lipgloss.NewStyle().Background(lipgloss.Color("#16001E")).Foreground(lipgloss.Color("#F7B2B7")),
		errorStyle:     lipgloss.NewStyle().Background(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#7F2982")),
		cursorStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFA00")),
	}

	var menu = &StartMenuView{
		cursor: 0,
		items:  []*MenuEntry{},
		style:  globalStyle,
	}

	exitView := &ExitView{}
	exitMenuItem := &MenuEntry{text: "Exit", view: exitView}

	var typingView = &TypingView{style: globalStyle, parentView: menu} //TODO would like to allow user-defined word count in this view.
	var typingMenuItem = &MenuEntry{text: "Typing Test", view: typingView}

	menu.items = append(menu.items, typingMenuItem)
	menu.items = append(menu.items, exitMenuItem)

	p := tea.NewProgram(menu)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
