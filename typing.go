package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TypeModel struct {
	cursor    int //TODO: maybe use the textinput bubble instead of dealing with a cursor like this...
	phrase    string
	userInput string
}

var (
	background         = lipgloss.Color("#16001E")
	phraseColor        = lipgloss.Color("#F7B2B7")
	errColor           = lipgloss.Color("#7F2982")
	errBackgroundColor = lipgloss.Color("#000000")
	inputColor         = lipgloss.Color("#DE639A")

	containerStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingTop(1).
			PaddingLeft(2)

	inputStyle  = lipgloss.NewStyle().Background(background).Foreground(inputColor).Render // a direct reference to the function, not an invocation?
	phraseStyle = lipgloss.NewStyle().Background(background).Foreground(phraseColor).Render
	errorStyle  = lipgloss.NewStyle().Background(errBackgroundColor).Foreground(errColor).Render
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
		case tea.KeyCtrlH: //ctrl + backspace
			// find the most recent space character before the current cursor position
			// set cursor to that index
			// delete user phrase after index.
		}

	case tea.WindowSizeMsg:
		containerStyle = containerStyle.Width(msg.Width)
	}
	return model, nil
}

func (model TypeModel) View() string {
	totalLength := len(model.phrase)
	if totalLength < len(model.userInput) {
		totalLength = len(model.userInput)
	}

	doc := ""
	for i := 0; i < totalLength; i++ {
		switch {
		//no input - always phrase
		case len(model.userInput) == 0:
			doc += phraseStyle(string(model.phrase[i]))
		// input too long, always error
		case i > len(model.phrase)-1:
			doc += errorStyle(string(model.userInput[i]))
		//input too short, always phrase
		case i > len(model.userInput)-1:
			doc += phraseStyle(string(model.phrase[i]))
		//match
		case model.userInput[i] == model.phrase[i]:
			doc += inputStyle(string(model.userInput[i]))
		//nomatch
		case model.userInput[i] != model.phrase[i]:
			doc += errorStyle(string(model.userInput[i]))
		default:
			panic("view render unreachable statement")
		}
	}

	return containerStyle.Render(doc)
}

func initialModel() TypeModel {
	return TypeModel{
		cursor:    0,
		phrase:    "sample text here.",
		userInput: "",
	}
}
