package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type ViewStyle struct {
	containerStyle lipgloss.Style
	inputStyle     lipgloss.Style
	phraseStyle    lipgloss.Style
	errorStyle     lipgloss.Style
	cursorStyle    lipgloss.Style
}

func main() {

	settingsFile := LoadUserSettings()
	appStyle := createStyle(settingsFile)
	wordCount := settingsFile.WordCount

	var menu = &StartMenuView{
		cursor: 0,
		items:  []*MenuEntry{},
		style:  appStyle,
	}

	exitView := &ExitView{}
	exitMenuItem := &MenuEntry{text: "Exit", view: exitView}

	typingView := &TypingView{style: appStyle, parentView: menu, wordCount: wordCount}
	typingMenuItem := &MenuEntry{text: "Typing Test", view: typingView}

	optionsView := &OptionsView{style: appStyle, parentView: menu, settings: &settingsFile}
	optionsMenuItem := &MenuEntry{text: "Options", view: optionsView}

	menu.items = append(menu.items, typingMenuItem)
	menu.items = append(menu.items, exitMenuItem)
	menu.items = append(menu.items, optionsMenuItem)

	p := tea.NewProgram(menu)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func createStyle(settingsFile SettingsFile) ViewStyle {
	generalBackground := lipgloss.Color(settingsFile.Style.Background)
	errorBackground := lipgloss.Color(settingsFile.Style.Errorbackground)
	cursorColor := lipgloss.Color(settingsFile.Style.Cursor)
	untypedTextColor := lipgloss.Color(settingsFile.Style.Text)
	correctTextColor := lipgloss.Color(settingsFile.Style.Correct)
	errorTextColor := lipgloss.Color(settingsFile.Style.Err)

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

	return globalStyle
}
