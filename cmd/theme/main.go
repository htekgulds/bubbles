package theme

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	base       lipgloss.Style
	border     lipgloss.TerminalColor
	background lipgloss.TerminalColor
	highlight  lipgloss.TerminalColor
	brand      lipgloss.TerminalColor
	error      lipgloss.TerminalColor
	body       lipgloss.TerminalColor
	accent     lipgloss.TerminalColor
}

func New() Theme {
	base := Theme{}
	base.background = lipgloss.AdaptiveColor{Dark: "#000000", Light: "#FBFCFD"}
	base.border = lipgloss.AdaptiveColor{Dark: "#3A3F42", Light: "#D7DBDF"}
	base.body = lipgloss.AdaptiveColor{Dark: "#889096", Light: "#889096"}
	base.accent = lipgloss.AdaptiveColor{Dark: "#FFFFFF", Light: "#11181C"}
	base.brand = lipgloss.Color("#FF5C00")
	base.error = lipgloss.Color("203")
	base.highlight = base.brand

	base.base = lipgloss.NewStyle().Foreground(base.body)
	return base
}

func (b Theme) Body() lipgloss.TerminalColor {
	return b.body
}

func (b Theme) Highlight() lipgloss.TerminalColor {
	return b.highlight
}

func (b Theme) Brand() lipgloss.TerminalColor {
	return b.brand
}

func (b Theme) Background() lipgloss.TerminalColor {
	return b.background
}

func (b Theme) Accent() lipgloss.TerminalColor {
	return b.accent
}

func (b Theme) Base() lipgloss.Style {
	return b.base
}

func (b Theme) TextBody() lipgloss.Style {
	return b.Base().Foreground(b.body)
}

func (b Theme) TextAccent() lipgloss.Style {
	return b.Base().Foreground(b.accent)
}

func (b Theme) TextHighlight() lipgloss.Style {
	return b.Base().Foreground(b.highlight)
}

func (b Theme) TextBrand() lipgloss.Style {
	return b.Base().Foreground(b.brand)
}

func (b Theme) TextError() lipgloss.Style {
	return b.Base().Foreground(b.error)
}

func (b Theme) PanelError() lipgloss.Style {
	return b.Base().Background(b.error).Foreground(b.accent)
}

func (b Theme) Border() lipgloss.TerminalColor {
	return b.border
}
