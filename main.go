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
	generalBackground := lipgloss.Color("#558B6A")
	errorBackground := lipgloss.Color("#ED6B86")
	cursorColor := lipgloss.Color("#D58936")
	untypedTextColor := lipgloss.Color("#F7B2B7")
	correctTextColor := lipgloss.Color("#DE639A")
	errorTextColor := lipgloss.Color("#7F2982")

	var globalStyle = ViewStyle{
		containerStyle: lipgloss.NewStyle().Background(generalBackground).
			PaddingBottom(1).
			PaddingLeft(1).
			PaddingTop(1).
			MarginLeft(1),

		inputStyle:  lipgloss.NewStyle().Foreground(correctTextColor).Background(generalBackground),
		phraseStyle: lipgloss.NewStyle().Foreground(untypedTextColor).Background(generalBackground),
		errorStyle:  lipgloss.NewStyle().Foreground(errorTextColor).Background(errorBackground),
		cursorStyle: lipgloss.NewStyle().Foreground(cursorColor).Background(generalBackground),
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
