package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/gitty/vcs"
)

type Model struct {
	issues   []vcs.Issue
	prs      []vcs.PullRequest
	branches []vcs.Branch
	commits  []vcs.Commit
	KeyMap   KeyMap
}

func NewModel() Model {
	// TODO: populate me senpai
	return Model{
		KeyMap: KeyMapSetup(),
	}
}

type KeyMap struct {
	Quit key.Binding
}

func KeyMapSetup() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.KeyMap.Quit) {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return parseRepository()
}
