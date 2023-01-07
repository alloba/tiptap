package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// The TypingView is the main user-experience struct.
// This stores state and function definitions for the actual typing test that a user takes.
type TypingView struct {
	phrase     string
	userInput  string
	style      ViewStyle
	parentView tea.Model
	startTime  time.Time
	wordCount  int
}

func (model *TypingView) Init() tea.Cmd {
	model.phrase = GenerateTypingPhrase(model.wordCount)
	model.userInput = ""
	model.startTime = time.Now()
	return nil
}

func (model *TypingView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// tracking these explicitly allows for easier checking of test completion
	var returnModel = model
	var returnCmd tea.Cmd = nil

	//event processing
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		//handling exit case separately to allow for screen wipe.
		case tea.KeyCtrlC, tea.KeyEscape:
			return &ExitView{}, tea.Quit
		default:
			returnModel, returnCmd = processUserInputEvents(msg, model)
		}
	case tea.WindowSizeMsg:
		model.style.containerStyle = model.style.containerStyle.Width(msg.Width)
	}

	//has the typing test been completed?
	if len(model.userInput) == len(model.phrase) {
		return &ResultsView{
				parentView:   model.parentView,
				style:        model.style,
				elapsedTime:  time.Now().Sub(model.startTime),
				targetPhrase: model.phrase,
				actualPhrase: model.userInput,
			},
			nil
	}
	return returnModel, returnCmd
}

// The View function travels the length of the user input and generated phrase to apply style to each character.
// Different styling is applied based on whether the current character is unmodified, correct, or in error.
// I expect the finer details of this to change over time (which is why there is a weird child function being called)
func (model *TypingView) View() string {
	totalLength := len(model.phrase)
	if totalLength < len(model.userInput) {
		totalLength = len(model.userInput)
	}

	doc := ""
	for i := 0; i < totalLength; i++ {
		doc += renderTechnique_errorPriority(*model, i)
	}
	doc += "\n"
	elapsedTimeSeconds := math.Floor(time.Now().Sub(model.startTime).Seconds())
	doc += fmt.Sprintf("%v", elapsedTimeSeconds) //TODO: it would be nice if this updated live. but currently it's on whatever standard render loop.

	return model.style.containerStyle.Render(doc)
}

// Handle all user input events during the update loop.
func processUserInputEvents(msg tea.KeyMsg, model *TypingView) (*TypingView, tea.Cmd) {
	switch msg.Type {

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
func renderTechnique_errorPriority(model TypingView, i int) string {

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
		// if there is no character, the foreground will not be rendered. so have to replace with a block char
		if string(model.phrase[i]) == " " {
			return model.style.cursorStyle.Render("█")
		} else {
			return model.style.cursorStyle.Render(string(model.phrase[i]))
		}

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
