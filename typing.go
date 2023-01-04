package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TypeModel struct {
	phrase    string
	userInput string
	style     Style
}

func NewTypingModel() TypeModel {
	style := Style{
		containerStyle: lipgloss.NewStyle().Bold(true).PaddingTop(1).PaddingLeft(2),
		inputStyle:     lipgloss.NewStyle().Background(lipgloss.Color("#16001E")).Foreground(lipgloss.Color("#DE639A")),
		phraseStyle:    lipgloss.NewStyle().Background(lipgloss.Color("#16001E")).Foreground(lipgloss.Color("#F7B2B7")),
		errorStyle:     lipgloss.NewStyle().Background(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#7F2982")),
		cursorStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFA00")),
	}

	return TypeModel{
		phrase:    GenerateTypingPhrase(20),
		userInput: "",
		style:     style,
	}
}

func (model TypeModel) Init() tea.Cmd {
	return nil
}

func (model TypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// tracking these explicitly allows for easier checking of test completion
	var returnModel TypeModel = model
	var returnCmd tea.Cmd = nil

	//event processing
	switch msg := msg.(type) {
	case tea.KeyMsg:
		returnModel, returnCmd = processUserInputEvents(msg, model)

	case tea.WindowSizeMsg:
		model.style.containerStyle = model.style.containerStyle.Width(msg.Width)
	}

	//has the typing test been completed?
	if len(model.userInput) == len(model.phrase)-1 {
		return model, tea.Quit
	}
	return returnModel, returnCmd
}

// Rendering the typing view is done in two major parts.
// First, each character on the screen is processed and has style rules applied to it based on
//   whether it is the cursor position, untyped, typed correct, or typed incorrect character.
// Second, the entire string is wrapped in a container that defines things like max width and spacing.
func (model TypeModel) View() string {
	totalLength := len(model.phrase)
	if totalLength < len(model.userInput) {
		totalLength = len(model.userInput)
	}

	doc := ""
	for i := 0; i < totalLength; i++ {
		doc += renderTechnique_errorPriority(model, i)
	}

	return model.style.containerStyle.Render(doc)
}

// Handle all user input events during the update loop.
func processUserInputEvents(msg tea.KeyMsg, model TypeModel) (TypeModel, tea.Cmd) {
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

// Incorrectly typed characters overwrite target characters in the view.
// This makes it easy to see what exactly was typed, but makes it harder to recover quickly.
func renderTechnique_errorPriority(model TypeModel, i int) string {

	switch {
	//no user input - color the first character as the cursor, phrase style otherwise.
	case len(model.userInput) == 0:
		if i == 0 {
			return model.style.cursorStyle.Render(string(model.phrase[i]))
		} else {
			return model.style.phraseStyle.Render(string(model.phrase[i]))
		}

	// cursor position - apply cursor style.
	case i == len(model.userInput):
		return model.style.cursorStyle.Render(string(model.phrase[i]))

	// input too long, always error
	case i > len(model.phrase)-1:
		return model.style.errorStyle.Render(string(model.userInput[i]))

	//input too short, always phrase
	case i > len(model.userInput)-1:
		return model.style.phraseStyle.Render(string(model.phrase[i]))

	//match
	case model.userInput[i] == model.phrase[i]:
		return model.style.inputStyle.Render(string(model.userInput[i]))

	//nomatch
	case model.userInput[i] != model.phrase[i]:
		return model.style.errorStyle.Render(string(model.userInput[i]))

	default:
		panic("view render unreachable statement")
	}
}
