package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

// DaisyUI Dark Theme

// Colors
var (
	base100  = lipgloss.Color("#1d232a")
	base200  = lipgloss.Color("#191e24")
	base300  = lipgloss.Color("#15191e")
	baseText = lipgloss.Color("#ecf9ff")

	primary       = lipgloss.Color("#605dff")
	primaryText   = lipgloss.Color("#edf1fe")
	secondary     = lipgloss.Color("#f43098")
	secondaryText = lipgloss.Color("#f9e4f0")

	accent      = lipgloss.Color("#00d3bb")
	accentText  = lipgloss.Color("#084d49")
	neutral     = lipgloss.Color("#09090b")
	neutralText = lipgloss.Color("#e4e4e7")

	info        = lipgloss.Color("#00bafe")
	infoText    = lipgloss.Color("#042e49")
	success     = lipgloss.Color("#00d390")
	successText = lipgloss.Color("#004c39")

	warning     = lipgloss.Color("#fcb700")
	warningText = lipgloss.Color("#793205")
	danger      = lipgloss.Color("#ff627d")
	dangerText  = lipgloss.Color("#4d0218")
)

var (
	base100light  = lipgloss.Color("#ffffff")
	base200light  = lipgloss.Color("#f8f8f8")
	base300light  = lipgloss.Color("#eeeeee")
	baseTextlight = lipgloss.Color("#18181b")

	primarylight       = lipgloss.Color("#422ad5")
	primaryTextlight   = lipgloss.Color("#e0e7ff")
	secondarylight     = lipgloss.Color("#f43098")
	secondaryTextlight = lipgloss.Color("#f9e4f0")
)

var baseStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{
		Light: string(baseTextlight),
		Dark:  string(baseText),
	})

var (
	Text = baseStyle
	// Hyperlink
	Link = lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Foreground(lipgloss.AdaptiveColor{
			Light: string(primarylight),
			Dark:  string(primary),
		})

	// Header
	Header = baseStyle.
		Padding(0, 1).
		Bold(true).
		Foreground(lipgloss.AdaptiveColor{
			Light: string(primaryTextlight),
			Dark:  string(primaryText),
		}).
		Background(lipgloss.AdaptiveColor{
			Light: string(primarylight),
			Dark:  string(primary),
		})

	// Highlight
	Mark = baseStyle.
		Padding(0, 1).
		Foreground(warningText).
		Background(warning)

	// Italic
	I = baseStyle.Italic(true)

	// Underline
	U = baseStyle.Underline(true)

	// Bold
	B = baseStyle.Bold(true)

	// Strikethrough
	S = baseStyle.Strikethrough(true)

	// Tick
	Tick = baseStyle.
		Foreground(successText).
		Render("✓")

	// Cross
	Cross = baseStyle.
		Foreground(dangerText).
		Render("✕")

	// Bang
	Bang = baseStyle.
		Foreground(warningText).
		Render("!")

	// Text Colors
	Base = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{
			Light: string(baseTextlight),
			Dark:  string(baseText),
		})

	Primary = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{
			Light: string(primaryTextlight),
			Dark:  string(primaryText),
		})

	Secondary = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: string(secondaryTextlight),
			Dark:  string(secondaryText),
		})

	Info    = lipgloss.NewStyle().Foreground(infoText)
	Success = lipgloss.NewStyle().Foreground(successText)
	Warning = lipgloss.NewStyle().Foreground(warningText)
	Danger  = lipgloss.NewStyle().Foreground(dangerText)
	Accent  = lipgloss.NewStyle().Foreground(accentText)
	Neutral = lipgloss.NewStyle().Foreground(neutralText)

	// Background Colors
	Base100Background = Base.Background(lipgloss.AdaptiveColor{
		Light: string(base100light),
		Dark:  string(base100),
	})
	Base200Background = Base.Background(lipgloss.AdaptiveColor{
		Light: string(base200light),
		Dark:  string(base200),
	})
	Base300Background = Base.Background(lipgloss.AdaptiveColor{
		Light: string(base300light),
		Dark:  string(base300),
	})
	PrimaryBackground = Primary.Background(lipgloss.AdaptiveColor{
		Light: string(primarylight),
		Dark:  string(primary),
	})
	SecondaryBackground = Secondary.Background(lipgloss.AdaptiveColor{
		Light: string(secondarylight),
		Dark:  string(secondary),
	})
	InfoBackground    = Info.Background(info)
	SuccessBackground = Success.Background(success)
	WarningBackground = Warning.Background(warning)
	DangerBackground  = Danger.Background(danger)
	AccentBackground  = Accent.Background(accent)
	NeutralBackground = Neutral.Background(neutral)

	// Logging defines a PurpleClay themed [logging] style that supports both light and
	// dark terminals
	//
	// [logging]: https://github.com/charmbracelet/log
	Logging = &log.Styles{
		Timestamp: baseStyle,
		Caller:    baseStyle.Faint(true),
		Prefix:    baseStyle.Bold(true).Faint(true),
		Message:   baseStyle.MarginRight(2),
		Key: baseStyle.
			Foreground(lipgloss.AdaptiveColor{
				Light: string(primarylight),
				Dark:  string(primary),
			}),
		Value:     baseStyle,
		Separator: baseStyle.Faint(true),
		Levels: map[log.Level]lipgloss.Style{
			log.DebugLevel: baseStyle.
				SetString(Tick).
				Bold(true).
				MaxWidth(2),
			log.InfoLevel: baseStyle.
				SetString(Tick).
				Bold(true).
				MaxWidth(2),
			log.WarnLevel: baseStyle.
				SetString(Bang).
				Bold(true).
				MaxWidth(2),
			log.ErrorLevel: baseStyle.
				SetString(Cross).
				Bold(true).
				MaxWidth(2),
			log.FatalLevel: baseStyle.
				SetString(Cross).
				Bold(true).
				MaxWidth(2),
		},
		Keys: map[string]lipgloss.Style{
			"err":   baseStyle.Foreground(danger),
			"error": baseStyle.Foreground(danger),
		},
		Values: map[string]lipgloss.Style{},
	}
)
