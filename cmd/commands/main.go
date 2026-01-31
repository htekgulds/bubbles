package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/htekgulds/bubbles/cmd/theme"
	"github.com/spf13/cobra"
)

type weatherMsg *WeatherData

type model struct {
	weather *WeatherData
	err     error
	width   int
	height  int
	loading bool
}

func (m model) Init() tea.Cmd {
	return fetchWeather("Ankara")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("Update", "msg", msg)
	switch _msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = _msg.Width
		m.height = _msg.Height
		return m, nil
	case tea.KeyMsg:
		switch _msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case weatherMsg:
		m.weather = _msg
		m.err = nil
		return m, nil
	case error:
		m.err = _msg
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		errorStyle := theme.Danger.
			Align(lipgloss.Center).
			Padding(2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.AdaptiveColor{
				Light: "#4d0218",
				Dark:  "#ff627d",
			})
		return errorStyle.Render("❌ Error: " + m.err.Error())
	}

	if m.weather == nil {
		loadingStyle := theme.Info.
			Align(lipgloss.Center).
			Padding(2)
		return loadingStyle.Render("🌤️  Hava durumu bilgisi yükleniyor...")
	}

	// Get weather emoji based on condition
	weatherEmoji := getWeatherEmoji(m.weather.Condition)

	// Card container style
	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.AdaptiveColor{
			Light: "#422ad5",
			Dark:  "#605dff",
		}).
		Padding(1, 2).
		Margin(1).
		Width(50).
		Align(lipgloss.Center)

	// Header style (location)
	headerStyle := theme.Header.
		Width(46).
		Align(lipgloss.Center).
		MarginBottom(1)

	// Temperature style (large and prominent)
	tempStyle := theme.Accent.
		Bold(true).
		Align(lipgloss.Center).
		Margin(1, 0)

	// Condition style
	conditionStyle := theme.Base.
		Italic(true).
		Align(lipgloss.Center).
		MarginBottom(1)

	// Info box style
	infoBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.AdaptiveColor{
			Light: "#042e49",
			Dark:  "#00bafe",
		}).
		Padding(0, 1).
		Margin(0, 1)

	// Info label style
	infoLabelStyle := theme.Info.
		Bold(true).
		MarginRight(1)

	// Info value style
	infoValueStyle := theme.Base

	// Build the UI
	header := headerStyle.Render("📍 " + m.weather.Location)

	temperature := tempStyle.Render(weatherEmoji + " " + m.weather.Temperature)

	condition := conditionStyle.Render(m.weather.Condition)

	// Info section
	humidityBox := infoBoxStyle.Render(
		infoLabelStyle.Render("💧 Humidity:") +
			infoValueStyle.Render(m.weather.Humidity),
	)

	windBox := infoBoxStyle.Render(
		infoLabelStyle.Render("💨 Wind:") +
			infoValueStyle.Render(m.weather.WindSpeed),
	)

	infoRow := lipgloss.JoinHorizontal(lipgloss.Center, humidityBox, windBox)

	// Combine all elements
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		temperature,
		condition,
		"",
		infoRow,
	)

	// Center the card on screen
	if m.width > 0 {
		cardStyle = cardStyle.Width(min(50, m.width-4))
		content = lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			cardStyle.Render(content),
		)
	} else {
		content = cardStyle.Render(content)
	}

	// Footer with instructions
	footerStyle := theme.Neutral.
		Faint(true).
		Align(lipgloss.Center).
		MarginTop(1)

	footer := footerStyle.Render("Press 'q' or 'esc' to quit")

	return lipgloss.JoinVertical(lipgloss.Center, content, footer)
}

func getWeatherEmoji(condition string) string {
	condition = strings.ToLower(condition)
	switch {
	case strings.Contains(condition, "sun") || strings.Contains(condition, "clear"):
		return "☀️"
	case strings.Contains(condition, "cloud"):
		return "☁️"
	case strings.Contains(condition, "rain"):
		return "🌧️"
	case strings.Contains(condition, "storm") || strings.Contains(condition, "thunder"):
		return "⛈️"
	case strings.Contains(condition, "snow"):
		return "❄️"
	case strings.Contains(condition, "fog") || strings.Contains(condition, "mist"):
		return "🌫️"
	case strings.Contains(condition, "wind"):
		return "💨"
	default:
		return "🌤️"
	}
}

func fetchWeather(location string) tea.Cmd {
	return func() tea.Msg {
		weather, err := FetchWeather(location)
		if err != nil {
			return fmt.Errorf("failed to fetch weather: %w", err)
		}
		return weatherMsg(weather)
	}
}

func initialModel() tea.Model {
	return model{loading: true}
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "Run command example (weather api sample)",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(commandsCmd)
}
