package main

import tea "github.com/charmbracelet/bubbletea"

// The ExitView is a bubbletea view struct that only serves the purpose of exiting the application.
// Wrapping the Quit command within a struct allows it to better fit into the overall application flow that I'm going for.
// It also gives me an easy way to ensure a screen wipe happens on exit. 
// (Since I couldn't find a nice way to integrate the clear screen fuction into the code.)
type ExitView struct {
}

func (model *ExitView) Init() tea.Cmd {
	return tea.Quit
}

func (model *ExitView) View() string {
	return ""
}

func (model *ExitView) Update(a tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}
