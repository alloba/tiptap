package main

import tea "github.com/charmbracelet/bubbletea"

// The StartMenuView is the entry point for user interaction.
// It contains any directly connected views, and implements a simple menu to allow navigation.
type StartMenuView struct {
	cursor int
	items  []*MenuEntry
	style  ViewStyle
}

type MenuEntry struct {
	text string
	view tea.Model
}

func (model *StartMenuView) Init() tea.Cmd {
	return nil
}

// The Update function processes user inputs as the user navigates the menu.
// There should be nothing surprising in here, except for one thing -
// When the user selects a menu option, the selected View's Init() function is invoked.
// Meaning, any initialization that the child view needs should be defined there.
func (model *StartMenuView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "ctrl+c", "q", "esc":
			return &ExitView{}, tea.Quit
		case "enter", "space":
			cmd := model.items[model.cursor].view.Init()
			return model.items[model.cursor].view, cmd
		}

	case tea.WindowSizeMsg: // NOTE* On Windows OS, this will only fire once on initial load. Currently it doesn't support the SIGWINCH signal.
		model.style.containerStyle = model.style.containerStyle.Width(msg.Width)
	}

	return model, nil
}

func (model *StartMenuView) View() string {
	docString := ""
	for i, item := range model.items {
		append := "\n" //this is awful but i am lazy. dont add an newline to the render if it's the last item in the list.
		if i == len(model.items)-1 {
			append = ""
		}
		if model.cursor == i {
			docString += model.style.cursorStyle.Render(item.text) + append
		} else {
			docString += model.style.phraseStyle.Render(item.text) + append
		}
	}
	return model.style.containerStyle.Render(docString)
}
