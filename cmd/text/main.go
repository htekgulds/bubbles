package text

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle   = focusedStyle
	focusedButton = focusedStyle.Render("[ Submit ]")
	buttonStyle   = fmt.Sprintf("%s %s", focusedButton, blurredStyle.Render("enter"))
)

type errMsg error

type model struct {
	input textinput.Model
	name  string
	err   error
}

func initialModel() model {
	i := textinput.New()
	i.Placeholder = "Enter your name"
	i.Focus()
	i.CharLimit = 20
	i.Width = 20
	i.Cursor.Style = cursorStyle
	i.PromptStyle = focusedStyle
	i.TextStyle = focusedStyle
	return model{input: i}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			m.name = m.input.Value()
			return m, nil
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	if m.name != "" {
		return fmt.Sprintf("Hello, %s!", m.name)
	}
	return fmt.Sprintf(
		"Write your name, please\n\n%s\n%s\n",
		m.input.View(),
		buttonStyle,
	)
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Run bubbles text input example",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(textCmd)
}
