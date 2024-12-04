package view

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle  = lipgloss.NewStyle().Align(lipgloss.Center).Background(lipgloss.Color(20))
	inputStyle   = lipgloss.NewStyle().Align(lipgloss.Left).Background(lipgloss.Color(40)).Border(lipgloss.RoundedBorder())
	contentStyle = lipgloss.NewStyle().Align(lipgloss.Left, lipgloss.Top)
)

type InputTest struct {
	Msg string
}

type model struct {
	isQuitting    bool
	width, height int
	table         table.Model
	input         textinput.Model
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
		v.Width = m.width/len(cols) - 2
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
	}

	return m, nil, preventPropergation
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case InputTest:
		m.table.SetRows(append(m.table.Rows(), table.Row{"tbd", msg.Msg}))
		m.table.GotoBottom()
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
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.isQuitting {
		return ""
	}

	header := headerStyle.Width(m.width).Render("header")
	input := inputStyle.Width(m.width - 2).Render(m.input.View())
	content := contentStyle.
		Width(m.width).
		Height(m.height - 4).
		Render(lipgloss.JoinHorizontal(lipgloss.Bottom, m.table.View()))

	return lipgloss.JoinVertical(lipgloss.Top, header, content, input)
}

func CreateRootProgram() *tea.Program {
	columns := []table.Column{
		{Title: "timestamp", Width: 4},
		{Title: "log", Width: 10},
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

	return tea.NewProgram(model{table: t, input: input})
}
