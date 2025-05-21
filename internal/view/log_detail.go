package view

import (
	"strings"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	detailContentStyle = lipgloss.NewStyle().Padding(0, 1)
)

type keyMap struct {
	copy  key.Binding
	close key.Binding
}

type logDetailModel struct {
	id            string
	keyMap        keyMap
	displayConfig config.DisplayConfig
	width, height int
	table         table.Model
	help          help.Model
}

func newLogDetail(id string, displayConfig config.DisplayConfig, window tea.WindowSizeMsg, repo aggregator.LogRepository) logDetailModel {
	columns := []table.Column{{Title: "", Width: 30}, {Title: "", Width: window.Width - 30}}
	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)
	s := table.DefaultStyles()
	s.Header = lipgloss.NewStyle()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("240")).
		Bold(false)
	t.SetStyles(s)

	log, err := repo.GetById(id)
	if err == nil {
		rows := getRowsByLogData(log.Data)
		if displayConfig.Detail.ShowRaw {
			rows = append(rows, table.Row{"raw", log.Raw})
		}
		t.SetRows(rows)
	}
	updateTableSize(&t, window)

	return logDetailModel{
		id:            id,
		keyMap:        defaultKeyMap(),
		displayConfig: displayConfig,
		table:         t,
		help:          help.New(),
		width:         window.Width,
		height:        window.Height,
	}
}

func (m logDetailModel) Init() tea.Cmd {
	return nil
}

func (m logDetailModel) updateSizes(msg tea.WindowSizeMsg) logDetailModel {
	m.width, m.height = msg.Width, msg.Height
	updateTableSize(&m.table, msg)
	return m
}

func updateTableSize(t *table.Model, window tea.WindowSizeMsg) {
	t.SetWidth(window.Width)
	t.SetHeight(window.Height - 5)
}

func getRowsByLogData(data aggregator.LogData) []table.Row {
	row := []table.Row{}
	for k, v := range data {
		row = append(row, table.Row{k, v})
	}
	return row
}

func (m logDetailModel) handleKeyPress(msg tea.KeyMsg) (logDetailModel, tea.Cmd, bool) {
	var cmd tea.Cmd
	var preventPropergation bool

	switch {
	case key.Matches(msg, m.keyMap.copy):
		value := m.table.SelectedRow()[1]
		_ = clipboard.WriteAll(value)
	case key.Matches(msg, m.keyMap.close):
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

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m logDetailModel) headerView() string {
	title := titleStyle.Render("Detail")
	line := strings.Repeat("─", max(0, m.width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m logDetailModel) footerView() string {
	line := strings.Repeat("─", max(0, m.width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func defaultKeyMap() keyMap {
	return keyMap{
		copy: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "copy to clipboard"),
		),
		close: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q/esc", "close"),
		),
	}
}

func (km keyMap) ShortHelp() []key.Binding {
	return []key.Binding{km.copy}
}

func (km keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{km.copy, km.close}}
}

func (m logDetailModel) View() string {
	return detailContentStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.headerView(),
			m.table.View(),
			m.footerView(),
			m.help.ShortHelpView(m.keyMap.ShortHelp()),
		),
	)
}
