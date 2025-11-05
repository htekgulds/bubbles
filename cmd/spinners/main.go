package spinners

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/htekgulds/bubbles/cmd/theme"
	"github.com/spf13/cobra"
)

type errMsg error

type model struct {
	spinners []spinner.Model
	quitting bool
	err      error
}

var themeStyles = theme.New()

func newSpinner(spinnerType spinner.Spinner) spinner.Model {
	s := spinner.New()
	s.Spinner = spinnerType
	s.Style = themeStyles.TextBrand()
	return s
}

func initialModel() model {
	ss := []spinner.Spinner{
		spinner.Dot,
		spinner.Ellipsis,
		spinner.Globe,
		spinner.Hamburger,
		spinner.Jump,
		spinner.Line,
		spinner.Moon,
		spinner.Pulse,
		spinner.Meter,
		spinner.MiniDot,
		spinner.Monkey,
		spinner.Points,
	}
	m := make([]spinner.Model, len(ss))
	for i, s := range ss {
		m[i] = newSpinner(s)
	}
	return model{spinners: m}
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

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(spinnerCmd)
}
