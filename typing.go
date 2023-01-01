package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TypeModel struct {
	phrase    string
	userInput string
}

var (
	background         = lipgloss.Color("#16001E")
	phraseColor        = lipgloss.Color("#F7B2B7")
	errColor           = lipgloss.Color("#7F2982")
	errBackgroundColor = lipgloss.Color("#FFFFFF")
	inputColor         = lipgloss.Color("#DE639A")

	cursorColor = lipgloss.Color("FAFA00")

	containerStyle = lipgloss.NewStyle().
			Bold(true).
			PaddingTop(1).
			PaddingLeft(2)

	// These are defined as functions to allow wrapping individual characters in the style easily.
	inputStyle  = lipgloss.NewStyle().Background(background).Foreground(inputColor).Render
	phraseStyle = lipgloss.NewStyle().Background(background).Foreground(phraseColor).Render
	errorStyle  = lipgloss.NewStyle().Background(errBackgroundColor).Foreground(errColor).Render
	cursorStyle = lipgloss.NewStyle().Foreground(cursorColor).Render
)

func (model TypeModel) Init() tea.Cmd {
	return nil
}

// Handle all user input events during the update loop.
func processUserInputEvents(msg tea.KeyMsg, model TypeModel) (tea.Model, tea.Cmd) {
	switch msg.Type {

	case tea.KeyCtrlC, tea.KeyEscape:
		return model, tea.Quit

	case tea.KeyRunes, tea.KeySpace:
		model.userInput += msg.String()

	case tea.KeyBackspace:
		if len(model.userInput) == 0 {
			return model, nil
		}
		model.userInput = model.userInput[0 : len(model.userInput)-1]

	case tea.KeyCtrlH: //ctrl + backspace
		if len(model.userInput) == 0 {
			return model, nil
		}

		var prevWordIndex = strings.LastIndex(model.userInput, " ")
		if prevWordIndex == -1 {
			model.userInput = "" // only one word exists, clear the entire field.
			return model, nil
		}

		model.userInput = model.userInput[0:prevWordIndex]
	}
	return model, nil
}

func (model TypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//has the typing test been completed?
	if len(model.userInput) == len(model.phrase) {
		print("Completed typing test")
		return model, tea.Quit
	}

	//event processing
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return processUserInputEvents(msg, model)

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
		//no user input - color the first character as the cursor, phrase style otherwise.
		case len(model.userInput) == 0:
			if i == 0 {
				doc += cursorStyle(string(model.phrase[i]))
			} else {
				doc += phraseStyle(string(model.phrase[i]))
			}

		// cursor position - apply cursor style.
		case i == len(model.userInput):
			doc += cursorStyle(string(model.phrase[i]))

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
		phrase:    GenerateTypingPhrase(50),
		userInput: "",
	}
}
