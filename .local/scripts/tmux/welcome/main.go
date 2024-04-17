package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var termHeight int
var termWidth int
var resultsFile string
var quote string

func loadsessions() ([]list.Item, error) {
	sessions, err := io.ReadAll(os.Stdin)
	if err != nil {
		return []list.Item{}, err
	}

	sess := strings.Split(string(sessions), "\n")
	items := []list.Item{}
	for _, s := range sess {
		cmps := strings.Split(s, " ")
		if len(cmps) < 3 {
			continue
		}

		wins, err := strconv.Atoi(cmps[1])
		if err != nil {
			return []list.Item{}, fmt.Errorf("Bad Session Date: %v", cmps[1])
		}

		sx := item{
			name:            cmps[0],
			index:           cmps[2],
			numberOfWindows: wins,
		}
		items = append(items, sx)
	}

	return items, nil
}

type model struct {
	sessions list.Model
	termH    int
	termW    int
	selected string
	banner   string
	quote    string
	isSSH    bool
}

func initialModel() (model, error) {
	sessions, err := loadsessions()
	if err != nil {
		return model{}, err
	}

	l := list.New(sessions, itemDelegate{}, termWidth, 20)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.SetShowFilter(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)
	l.SetShowPagination(false)
	l.Styles.Title = titleStyle
	l.DisableQuitKeybindings()
	l.Paginator.KeyMap = paginator.KeyMap{}
	l.KeyMap = list.KeyMap{}
	l.KeyMap.Filter = key.NewBinding(key.WithKeys("/"))
	l.InfiniteScrolling = true

	ssh := os.Getenv("SSH_CLIENT")

	return model{
		sessions: l,
		termH:    termHeight,
		termW:    termWidth,
		banner:   ascii(),
		quote:    quote,
		isSSH:    ssh != "",
	}, nil
}

func (m *model) resize() {
	m.resizeWithWidth(m.termW)
}

func (m *model) resizeWithWidth(w int) {
	m.termW = w
	banner.PaddingLeft((w / 2) - (lipgloss.Width(m.banner) / 2))
	pagination.Width(w)
	sessionItem.Width(w / 4)
	title.Width(w)
	help.Width(w)
	quoteStyle.Width(w)
}

func (m model) Init() tea.Cmd {
	m.resize()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.sessions.SetWidth(msg.Width)
		m.resizeWithWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "d":
			if m.sessions.FilterState() != list.Filtering {
				saveResults("@detach")
				return m, tea.Quit
			}

		case "X":
			if m.sessions.FilterState() != list.Filtering {
				saveResults("@disconnect")
				return m, tea.Quit
			}

		case "j", "down":
			if m.sessions.FilterState() != list.Filtering {
				m.sessions.CursorDown()
			}

		case "k", "up":
			if m.sessions.FilterState() != list.Filtering {
				m.sessions.CursorUp()
			}

		case tea.KeyEscape.String():
			if m.sessions.FilterState() == list.Filtering {
				m.sessions.FilterInput.Blur()
				m.sessions.ResetFilter()
				return m, nil
			}

		case "ctrl+c":
			if m.sessions.FilterState() != list.Filtering {
				return m, tea.Quit
			}
			m.sessions.FilterInput.Blur()
			m.sessions.ResetFilter()
			return m, nil

		case "q":
			if m.sessions.FilterState() != list.Filtering {
				return m, tea.Quit
			}

		case "enter", " ":
			item, ok := m.sessions.SelectedItem().(item)
			if !ok {
				panic("Something went wrong while selecting item")
			}
			m.selected = item.index
			saveResults(m.selected)
			return m, tea.Quit
		}
	}

	m.sessions, cmd = m.sessions.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := strings.Builder{}

	s.WriteString(banner.Render(m.banner))
	s.WriteString("\n")
	s.WriteString(title.Render("Sessions"))
	s.WriteString("\n")

	sessions := lipgloss.PlaceHorizontal(m.termW, lipgloss.Center, m.sessions.View())
	s.WriteString(sessions)
	s.WriteString("\n")
	if len(m.sessions.Items()) > m.sessions.Height() && !m.sessions.Paginator.OnLastPage() {
		s.WriteString(pagination.Render(""))
	}

	if m.sessions.FilterValue() != "" {
		s.WriteString(banner.Render(fmt.Sprintf(" %s", m.sessions.FilterValue())))
		s.WriteString("\n")
	}

	if m.sessions.FilterValue() == "" {
		bar := "Exit (q)  Filter (/)  Detach (d)"
		if m.isSSH {
			bar = fmt.Sprintf("%s  Disconnect (X)", bar)
		}

		s.WriteString(help.Render(bar))
		s.WriteString("\n")
	}

	s.WriteString(quoteStyle.Render(m.quote))
	s.WriteString("\n")

	return s.String()
}

func main() {
	flag.IntVar(&termWidth, "width", 0, "terminal width")
	flag.IntVar(&termHeight, "height", 0, "terminal height")
	flag.StringVar(&resultsFile, "result", "", "results file")
	flag.StringVar(&quote, "quote", "howdy", "quote")
	flag.Parse()

	m, err := initialModel()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func saveResults(res string) {
	f, err := os.Create(resultsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(res)
	if err != nil {
		log.Fatal(err)
	}
}
