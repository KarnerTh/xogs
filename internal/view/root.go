package view

import (
	"fmt"
	"os"

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
		v.Width = m.width / 5
		cols[i] = v
	}

	m.table.SetColumns(cols)
	return m
}

func (m model) handleKeyPress(msg tea.KeyMsg, cmd tea.Cmd) (model, tea.Cmd) {
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
		}
		return m, cmd
	case "q":
		if !m.input.Focused() {
			m.isQuitting = true
			return m, tea.Quit
		}
	case "ctrl+c":
		m.isQuitting = true
		return m, tea.Quit
	case "enter":
		if m.input.Focused() {
			m.table.Focus()
			m.input.Blur()
		}
	}

	return m, cmd
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m = m.updateSizes(msg)
	case tea.KeyMsg:
		m, cmd = m.handleKeyPress(msg, cmd)
		if cmd != nil {
			return m, cmd
		}
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

func ShowRoot() {
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 20},
	}

	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"3", "Shanghai", "China", "28,516,904"},
		{"4", "Dhaka", "Bangladesh", "22,478,116"},
		{"5", "São Paulo", "Brazil", "22,429,800"},
		{"6", "Mexico City", "Mexico", "22,085,140"},
		{"7", "Cairo", "Egypt", "21,750,020"},
		{"8", "Beijing", "China", "21,333,332"},
		{"9", "Mumbai", "India", "20,961,472"},
		{"10", "Osaka", "Japan", "19,059,856"},
		{"11", "Chongqing", "China", "16,874,740"},
		{"12", "Karachi", "Pakistan", "16,839,950"},
		{"13", "Istanbul", "Turkey", "15,636,243"},
		{"14", "Kinshasa", "DR Congo", "15,628,085"},
		{"15", "Lagos", "Nigeria", "15,387,639"},
		{"16", "Buenos Aires", "Argentina", "15,369,919"},
		{"17", "Kolkata", "India", "15,133,888"},
		{"18", "Manila", "Philippines", "14,406,059"},
		{"19", "Tianjin", "China", "14,011,828"},
		{"20", "Guangzhou", "China", "13,964,637"},
		{"21", "Rio De Janeiro", "Brazil", "13,634,274"},
		{"22", "Lahore", "Pakistan", "13,541,764"},
		{"23", "Bangalore", "India", "13,193,035"},
		{"24", "Shenzhen", "China", "12,831,330"},
		{"25", "Moscow", "Russia", "12,640,818"},
		{"26", "Chennai", "India", "11,503,293"},
		{"27", "Bogota", "Colombia", "11,344,312"},
		{"28", "Paris", "France", "11,142,303"},
		{"29", "Jakarta", "Indonesia", "11,074,811"},
		{"30", "Lima", "Peru", "11,044,607"},
		{"31", "Bangkok", "Thailand", "10,899,698"},
		{"32", "Hyderabad", "India", "10,534,418"},
		{"33", "Seoul", "South Korea", "9,975,709"},
		{"34", "Nagoya", "Japan", "9,571,596"},
		{"35", "London", "United Kingdom", "9,540,576"},
		{"36", "Chengdu", "China", "9,478,521"},
		{"37", "Nanjing", "China", "9,429,381"},
		{"38", "Tehran", "Iran", "9,381,546"},
		{"39", "Ho Chi Minh City", "Vietnam", "9,077,158"},
		{"40", "Luanda", "Angola", "8,952,496"},
		{"41", "Wuhan", "China", "8,591,611"},
		{"42", "Xi An Shaanxi", "China", "8,537,646"},
		{"43", "Ahmedabad", "India", "8,450,228"},
		{"44", "Kuala Lumpur", "Malaysia", "8,419,566"},
		{"45", "New York City", "United States", "8,177,020"},
		{"46", "Hangzhou", "China", "8,044,878"},
		{"47", "Surat", "India", "7,784,276"},
		{"48", "Suzhou", "China", "7,764,499"},
		{"49", "Hong Kong", "Hong Kong", "7,643,256"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
		{"50", "Riyadh", "Saudi Arabia", "7,538,200"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
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

	m := model{table: t, input: input}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
