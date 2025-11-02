package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type errMsg error

type model struct {
	spinners []spinner.Model
	quitting bool
	err      error
}

func newSpinner(spinnerType spinner.Spinner, color string) spinner.Model {
	if color == "" {
		color = "205"
	}
	s := spinner.New()
	s.Spinner = spinnerType
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
	return s
}

func initialModel() model {
	d := newSpinner(spinner.Dot, "205")
	e := newSpinner(spinner.Ellipsis, "205")
	g := newSpinner(spinner.Globe, "205")
	h := newSpinner(spinner.Hamburger, "205")
	j := newSpinner(spinner.Jump, "205")
	l := newSpinner(spinner.Line, "205")
	m := newSpinner(spinner.Moon, "205")
	p := newSpinner(spinner.Pulse, "205")
	mt := newSpinner(spinner.Meter, "205")
	md := newSpinner(spinner.MiniDot, "205")
	mn := newSpinner(spinner.Monkey, "205")
	pt := newSpinner(spinner.Points, "205")
	return model{spinners: []spinner.Model{d, e, g, h, j, l, m, p, mt, md, mn, pt}}
}

func (m model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, s := range m.spinners {
		cmds = append(cmds, s.Tick)
	}
	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}
	case errMsg:
		m.err = msg
		return m, nil
	default:
		var cmds = []tea.Cmd{}
		for i, s := range m.spinners {
			spinner, cmd := s.Update(msg)
			m.spinners[i] = spinner
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	strs := make([]string, len(m.spinners))
	for i, s := range m.spinners {
		strs[i] = s.View()
	}
	str := strings.Join(strs, " ")
	if m.quitting {
		return str + "\n"
	}
	return str
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var spinnerCmd = &cobra.Command{
	Use:   "spinners",
	Short: "Run all bubbles spinner animations",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func init() {
	rootCmd.AddCommand(spinnerCmd)
}
