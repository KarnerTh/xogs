package view

import tea "github.com/charmbracelet/bubbletea"

type logDetailModel struct {
	width, height int
	id            string
}

func newLogDetail(id string) logDetailModel {
	return logDetailModel{
		id: id,
	}
}

func (m logDetailModel) Init() tea.Cmd {
	return nil
}

func (m logDetailModel) updateSizes(msg tea.WindowSizeMsg) logDetailModel {
	m.width, m.height = msg.Width, msg.Height
	return m
}

func (m logDetailModel) handleKeyPress(msg tea.KeyMsg) (logDetailModel, tea.Cmd, bool) {
	var cmd tea.Cmd
	var preventPropergation bool

	switch msg.String() {
	case "esc", "q":
		return m, func() tea.Msg { return pushPageMsg{pageIdx: logListPage} }, preventPropergation
	}

	return m, cmd, preventPropergation
}

func (m logDetailModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m = m.updateSizes(msg)
	case tea.KeyMsg:
		model, cmd, preventPropergation := m.handleKeyPress(msg)
		if cmd != nil || preventPropergation {
			return model, cmd
		}

		m = model
	}

	return m, cmd
}

func (m logDetailModel) View() string {
	return m.id
}
