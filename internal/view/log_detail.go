package view

import (
	"fmt"
	"strings"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b).Padding(0, 1)
	}()

	detailContentStyle = lipgloss.NewStyle().Padding(0, 1)
	cellStyle          = lipgloss.NewStyle().PaddingRight(2)
	oddRowStyle        = cellStyle.Foreground(lipgloss.Color("250"))
	evenRowStyle       = cellStyle.Foreground(lipgloss.Color("245"))
	keyColumMaxWidth   = 15

	headerHeight, footerHeight = 3, 3
	widthPadding               = 2
	verticalMarginHeight       = headerHeight + footerHeight
)

type logDetailModel struct {
	id            string
	width, height int
	viewport      viewport.Model
	table         *table.Table
}

func newLogDetail(id string, window tea.WindowSizeMsg, repo aggregator.LogRepository) logDetailModel {
	var detailTable *table.Table
	viewport := viewport.New(window.Width-widthPadding, window.Height-verticalMarginHeight)
	viewport.YPosition = headerHeight

	log, err := repo.GetById(id)
	if err != nil {
		viewport.SetContent(err.Error())
	} else {
		detailTable = getLogDetail(*log)
		viewport.SetContent(detailTable.Render())
	}

	return logDetailModel{
		id:       id,
		viewport: viewport,
		table:    detailTable,
	}
}

func getLogDetail(log aggregator.Log) *table.Table {
	rows := [][]string{
		{"id", log.Id},
	}

	for key, value := range log.Data {
		rows = append(rows, []string{key, value})
	}

	rows = append(rows, []string{"raw", log.Raw})

	t := table.New().
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row%2 == 0:
				if col == 0 {
					return evenRowStyle.MaxWidth(keyColumMaxWidth)
				}
				return evenRowStyle
			default:
				if col == 0 {
					return oddRowStyle.MaxWidth(keyColumMaxWidth)
				}
				return oddRowStyle
			}
		})

	return t
}

func (m logDetailModel) Init() tea.Cmd {
	return nil
}

func (m logDetailModel) updateSizes(msg tea.WindowSizeMsg) logDetailModel {
	m.width, m.height = msg.Width, msg.Height
	m.viewport.Width = msg.Width - widthPadding
	m.viewport.Height = msg.Height - verticalMarginHeight

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

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m logDetailModel) headerView() string {
	title := titleStyle.Render("Detail")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m logDetailModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m logDetailModel) View() string {
	return detailContentStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.headerView(),
			m.viewport.View(),
			m.footerView(),
		),
	)
}
