package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

// OptionsView is ultimately meant to allow changes to all persisted settings (in SettingsFile),
// but I am lazy and this project is coming to a close in my mind.
// So! this is a very basic version that just allows changing the amount of words in the test.
type OptionsView struct {
	parentView tea.Model
	style      ViewStyle
	settings   *SettingsFile

	cursor int
	items  []*wpmOption
}

type wpmOption struct {
	count    int
	selected bool
}

func (model *OptionsView) Init() tea.Cmd {
	model.items = []*wpmOption{
		{count: 10, selected: false},
		{count: 25, selected: false},
		{count: 50, selected: false},
		{count: 100, selected: false},
		{count: 150, selected: false},
	}

	for _, item := range model.items {
		if item.count == model.settings.WordCount {
			item.selected = true
		}
	}
	return nil
}
func (model *OptionsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if model.cursor <= 0 {
				return model, nil
			}
			model.cursor -= 1
		case "down", "j":
			if model.cursor >= len(model.items)-1 {
				return model, nil
			}
			model.cursor += 1
		case "ctrl+c":
			return &ExitView{}, tea.Quit
		case "q", "esc":
			return model.parentView, nil
		case "enter", " ":
			for _, item := range model.items {
				if item.selected {
					item.selected = false
				}
			}
			model.items[model.cursor].selected = true
			model.settings.WordCount = model.items[model.cursor].count
			SaveUserSettings(*model.settings)
			return model, nil
		}

	case tea.WindowSizeMsg: // NOTE* On Windows OS, this will only fire once on initial load. Currently it doesn't support the SIGWINCH signal.
		model.style.containerStyle = model.style.containerStyle.Width(msg.Width)
	}
	return model, nil
}
func (model *OptionsView) View() string {
	doc := ""
	for i := range model.items {
		content := ""
		if model.items[i].selected {
			content += "* "
		} else {
			content += "  "
		}
		content += fmt.Sprintf("%v", model.items[i].count)

		if i == model.cursor {
			doc += model.style.cursorStyle.Render(content)
		} else {
			doc += model.style.phraseStyle.Render(content)
		}
		doc += "\n"
	}

	return model.style.containerStyle.Render(doc)
}
