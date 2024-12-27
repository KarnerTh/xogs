package view

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/KarnerTh/xogs/internal/observer"
	tea "github.com/charmbracelet/bubbletea"
)

type page int

const (
	logListPage page = iota
	detailPage  page = iota
)

type rootModel struct {
	curPageIdx  page
	pages       []tea.Model
	selectedLog *aggregator.Log
	isQuitting  bool
	window      tea.WindowSizeMsg
	repo        aggregator.LogRepository
}

func (m rootModel) Init() tea.Cmd {
	return nil
}

func (m rootModel) handleKeyPress(msg tea.KeyMsg) (rootModel, tea.Cmd, bool) {
	var preventPropergation bool

	switch msg.String() {
	case "ctrl+c":
		m.isQuitting = true
		return m, tea.Quit, preventPropergation
	}

	return m, nil, preventPropergation
}

func (m rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case aggregator.Notification:
	case logSelectedMsg:
		m.pages[detailPage] = newLogDetail(msg.id, m.window, m.repo)
		return m, func() tea.Msg { return pushPageMsg{pageIdx: detailPage} }
	case pushPageMsg:
		m.curPageIdx = msg.pageIdx
		switch msg.pageIdx {
		case logListPage:
			// refresh log list as the page was not getting any events
			return m, func() tea.Msg { return refreshMsg{} }
		}
	case tea.WindowSizeMsg:
		m.window = msg
	case tea.KeyMsg:
		model, cmd, preventPropergation := m.handleKeyPress(msg)
		if cmd != nil || preventPropergation {
			return model, cmd
		}

		m = model
	}

	m.pages[m.curPageIdx], cmd = m.pages[m.curPageIdx].Update(msg)
	return m, cmd
}

func (m rootModel) View() string {
	if m.isQuitting {
		return ""
	}

	return m.pages[m.curPageIdx].View()
}

func CreateRootProgram(displayConfig config.DisplayConfig, filter observer.Publisher[string], repo aggregator.LogRepository) *tea.Program {
	return tea.NewProgram(rootModel{
		curPageIdx: logListPage,
		repo:       repo,
		pages: []tea.Model{
			newLogList(displayConfig, filter),
			newLogDetail("", tea.WindowSizeMsg{}, repo),
		},
	}, tea.WithAltScreen(), tea.WithMouseAllMotion())
}
