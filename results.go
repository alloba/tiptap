package main

import (
	"fmt"
	"math"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ResultsView struct {
	style        ViewStyle
	parentView   tea.Model
	elapsedTime  time.Duration
	targetPhrase string
	actualPhrase string
}

func (model *ResultsView) Init() tea.Cmd {
	return nil
}

func (model *ResultsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return model, tea.Quit
		case tea.KeyEscape, tea.KeyBackspace:
			return model.parentView, nil
		}

	case tea.WindowSizeMsg: // NOTE* On Windows OS, this will only fire once on initial load. Currently it doesn't support the SIGWINCH signal.
		model.style.containerStyle = model.style.containerStyle.Width(msg.Width)
	}
	return model, nil
}

func (model *ResultsView) View() string {
	docString := ""

	timeString := fmt.Sprintf("Elapsed Time: %v seconds", math.Floor(model.elapsedTime.Seconds()))
	docString += model.style.phraseStyle.Render(timeString)

	docString += model.style.phraseStyle.Render("\n")

	accuracyString := fmt.Sprintf("Accuracy %.1f%%", calculateInputAccuracy(model.targetPhrase, model.actualPhrase))
	docString += model.style.phraseStyle.Render(accuracyString)
	return model.style.containerStyle.Render(docString)
}

func calculateInputAccuracy(target string, actual string) float32 {
	totalLength := len(target)
	correct := 0
	for i := 0; i < totalLength; i++ {
		if target[i] == actual[i] {
			correct += 1
		}
	}

	return float32(correct) / float32(totalLength) * 100
}
