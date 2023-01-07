package main

import tea "github.com/charmbracelet/bubbletea"

type OptionsView struct {
	parentView tea.Model
	style      ViewStyle
	settings   *SettingsFile
}

//TODO: all of this. need to think about how to edit values in the menu...

func (model *OptionsView) Init() tea.Cmd {
	return nil
}
func (model *OptionsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}
func (model *OptionsView) View() string {
	return model.style.containerStyle.Render(model.style.phraseStyle.Render("TODO"))
}
