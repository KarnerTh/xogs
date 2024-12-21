package view

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/KarnerTh/xogs/internal/observer"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	inputStyle   = lipgloss.NewStyle().Align(lipgloss.Left).Border(lipgloss.RoundedBorder())
	contentStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Top).PaddingTop(1)
)

type model struct {
	displayConfig   config.DisplayConfig
	filterPublisher observer.Publisher[string]
	isQuitting      bool
	follow          bool
	width, height   int
	table           table.Model
	input           textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) updateSizes(msg tea.WindowSizeMsg) model {
	m.width, m.height = msg.Width, msg.Height
	m.table.SetWidth(m.width)
	m.table.SetHeight(m.height - 4)
	m.input.Width = m.width - 10

	var cols = m.table.Columns()
	for i, v := range cols {
		v.Width = int(float32(m.width)*m.displayConfig.Columns[i].Width - 2)
		cols[i] = v
	}

	m.table.SetColumns(cols)
	return m
}

func (m model) handleKeyPress(msg tea.KeyMsg) (model, tea.Cmd, bool) {
	var preventPropergation bool

	switch msg.String() {
	case "esc":
		if m.input.Focused() {
			m.table.Focus()
			m.input.Blur()
		}
	case "i":
		if m.table.Focused() {
			m.table.Blur()
			m.input.Focus()
			preventPropergation = true
			return m, nil, preventPropergation
		}
	case "q":
		if !m.input.Focused() {
			m.isQuitting = true
			return m, tea.Quit, preventPropergation
		}
	case "ctrl+c":
		m.isQuitting = true
		return m, tea.Quit, preventPropergation
	case "enter":
		if m.input.Focused() {
			m.table.Focus()
			m.input.Blur()
		}
	case "up", "k":
		m.follow = false
	case "end", "G":
		m.follow = true
	}

	return m, nil, preventPropergation
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case aggregator.Notification:
		if msg.NewEntry != nil {
			row := mapLogToRow(m.displayConfig, *msg.NewEntry)
			m.table.SetRows(append(m.table.Rows(), row))
		} else if msg.BaseList != nil {
			rows := make([]table.Row, len(msg.BaseList))
			for i, v := range msg.BaseList {
				rows[i] = mapLogToRow(m.displayConfig, v)
			}
			m.table.SetRows(rows)
			m.table.GotoBottom()
		}

		if m.follow {
			m.table.GotoBottom()
		}
	case tea.WindowSizeMsg:
		m = m.updateSizes(msg)
	case tea.KeyMsg:
		model, cmd, preventPropergation := m.handleKeyPress(msg)
		if cmd != nil || preventPropergation {
			return model, cmd
		}

		m = model
	}

	input, cmd := m.input.Update(msg)
	if input.Value() != m.input.Value() {
		m.filterPublisher.Publish(input.Value())
	}
	m.input = input

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func mapLogToRow(displayConfig config.DisplayConfig, log aggregator.Log) table.Row {
	row := make(table.Row, len(displayConfig.Columns))
	for i, v := range displayConfig.Columns {
		if v.ValueKey == config.ValueKeyRaw {
			row[i] = log.Raw
		} else {
			row[i] = log.GetStringData(v.ValueKey)
		}
	}

	return row
}

func (m model) View() string {
	if m.isQuitting {
		return ""
	}

	input := inputStyle.Width(m.width - 2).Render(m.input.View())
	content := contentStyle.
		Width(m.width).
		Height(m.height - 3).
		Render(lipgloss.JoinHorizontal(lipgloss.Bottom, m.table.View()))

	return lipgloss.JoinVertical(lipgloss.Top, content, input)
}

func CreateRootProgram(displayConfig config.DisplayConfig, filter observer.Publisher[string]) *tea.Program {
	columns := make([]table.Column, len(displayConfig.Columns))
	for i, v := range displayConfig.Columns {
		columns[i] = table.Column{
			Title: v.Title,
			Width: 1,
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	input := textinput.New()
	input.Placeholder = "search and filter"

	return tea.NewProgram(model{
		displayConfig:   displayConfig,
		table:           t,
		input:           input,
		follow:          true,
		filterPublisher: filter,
	})
}
