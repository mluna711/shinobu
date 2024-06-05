package queueview

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type itemDelegate struct {
	height int
}

func (d itemDelegate) Height() int  { return d.height }
func (d itemDelegate) Spacing() int { return 0 }
func (d itemDelegate) Update(_ tea.Msg, m *list.Model) tea.Cmd {
	itemStyle.Width(m.Width())
	selectedItemStyle.Width(m.Width())
	return nil
}
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}
	isSelected := index == m.Index()
	bg := iconStyle.GetBackground()
	if isSelected {
		bg = selectedItemStyle.GetBackground()
	}

	fn := itemStyle.Render
	if isSelected {
		fn = selectedItemStyle.Render
	}

	icon := strings.Trim(i.Ascii, "\n")
	if icon == "" {
		icon = `//////////
//////////
//////////
//////////
//////////`
	}

	details := lipgloss.JoinVertical(lipgloss.Left, fn(fmt.Sprintf(" %s", i.Name)), fn(fmt.Sprintf(" %s", i.Artist)), fn(fmt.Sprintf(" %s", i.Duration)))
	if i.IsPlaying {
		details = lipgloss.JoinVertical(lipgloss.Left, details, fn(" Playing..."))
	}

	if isSelected {
		details = lipgloss.Place(m.Width()-6, d.Height(), lipgloss.Top, lipgloss.Left, details, lipgloss.WithWhitespaceBackground(bg))
	} else {
		details = lipgloss.Place(m.Width()-6, d.Height(), lipgloss.Top, lipgloss.Left, details, lipgloss.WithWhitespaceBackground(lipgloss.NoColor{}))
	}

	str := lipgloss.JoinHorizontal(lipgloss.Top, icon, details)

	fmt.Fprint(w, zone.Mark(i.ID, str))
}
