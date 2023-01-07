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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return &ExitView{}, tea.Quit
		case tea.KeyEscape, tea.KeyBackspace:
			return model.parentView, nil
		}

	case tea.WindowSizeMsg: // NOTE* On Windows OS, this will only fire once on initial load. Currently it doesn't support the SIGWINCH signal.
		model.style.containerStyle = model.style.containerStyle.Width(msg.Width)
	}
	return model, nil
}
func (model *OptionsView) View() string {
	return model.style.containerStyle.Render(model.style.phraseStyle.Render("TODO"))
}
