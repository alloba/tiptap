package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TypeModel struct {
	cursor    int
	phrase    string
	userInput string
}

var (
	background  = lipgloss.Color("#16001E") //black/purple
	phraseColor = lipgloss.Color("#F7B2B7") //light pink
	inputColor  = lipgloss.Color("#DE639A") // china pink / pink-purple

	//cursorColor = lipgloss.Color("#000000") // white
	//errColor    = lipgloss.Color("#7F2982") // ultra red / pink-red

	style = lipgloss.NewStyle().
		Bold(true).
		Foreground(phraseColor).
		Background(background).
		PaddingTop(1).
		PaddingLeft(2).
		Width(22)
)

func (model TypeModel) Init() tea.Cmd {
	return nil
}

func (model TypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return model, tea.Quit
		case tea.KeyRunes, tea.KeySpace:
			model.userInput += msg.String()
			model.cursor += 1
		case tea.KeyBackspace:
			if model.cursor > 0 {
				model.cursor -= 1
				model.userInput = model.userInput[0 : len(model.userInput)-1]
			}
			//TODO CTRL+BACKSPACE to delete current word.
		}

	case tea.WindowSizeMsg:
		style = style.Width(msg.Width)
	}
	return model, nil
}

func (model TypeModel) View() string {
	// TODO: I'm not sure of how to change the foreground color for single characters.
	//	     Rendering a single character width's worth of stuff is strangely difficult, which is needed to swap styles...
	//		 they have this kind of working in their example so look into that i guess.
	// 		 https://github.com/charmbracelet/lipgloss/blob/master/example/main.go
	//[1;38;2;247;178;183;48;2;22;0;30mplease work eventually[0m
	//[1;38;2;247;178;183;48;2;22;0;30mp[0m[1;38;2;247;178;183;48;2;22;0;30ml[0m

	//38;                   true color setting.
	//ESC[1;34;{...}m   --> graphics mode for cell

	doc := ""
	totalLength := len(model.phrase)
	if totalLength < len(model.userInput) {
		totalLength = len(model.userInput)
	}

	for i := 0; i < totalLength; i++ {
		if i < len(model.userInput) {
			doc += string(model.userInput[i])
		} else {
			doc += string(model.phrase[i])
		}
	}

	return style.Render(doc)
}

func initialModel() TypeModel {
	return TypeModel{
		cursor:    0,
		phrase:    "sample text here.",
		userInput: "",
	}
}
