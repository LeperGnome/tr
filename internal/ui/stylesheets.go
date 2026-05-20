package ui

import "github.com/charmbracelet/lipgloss"

type Stylesheet struct {
	SelectedPath lipgloss.Style

	FinfoPermissions lipgloss.Style
	FinfoLastUpdated lipgloss.Style
	FinfoSize        lipgloss.Style
	FinfoSep         lipgloss.Style

	OperationBar      lipgloss.Style
	OperationBarInput lipgloss.Style

	ErrBar lipgloss.Style

	HelpMsg     lipgloss.Style
	HelpContent lipgloss.Style

	TreeRegularFileName lipgloss.Style
	TreeDirecotryName   lipgloss.Style
	TreeLinkName        lipgloss.Style
	TreeMarkedNode      lipgloss.Style
	TreeSelectionArrow  lipgloss.Style
	TreeIndent          lipgloss.Style
	TreeIndentSelected  lipgloss.Style

	PlainTextPreview lipgloss.Style
}

var SystemStylesheet = Stylesheet{
	SelectedPath: lipgloss.NewStyle().Foreground(lipgloss.Color("2")),

	FinfoPermissions: lipgloss.NewStyle().Foreground(lipgloss.Color("3")),
	FinfoLastUpdated: lipgloss.NewStyle().Foreground(lipgloss.Color("7")),
	FinfoSize:        lipgloss.NewStyle().Foreground(lipgloss.Color("7")),
	FinfoSep:         lipgloss.NewStyle().Foreground(lipgloss.Color("0")),

	OperationBar:      lipgloss.NewStyle().Foreground(lipgloss.Color("7")),
	OperationBarInput: lipgloss.NewStyle().Background(lipgloss.Color("8")),

	ErrBar:      lipgloss.NewStyle().Foreground(lipgloss.Color("1")),
	HelpMsg:     lipgloss.NewStyle().Foreground(lipgloss.Color("3")),
	HelpContent: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),

	TreeRegularFileName: lipgloss.NewStyle().Foreground(lipgloss.Color("7")),
	TreeDirecotryName:   lipgloss.NewStyle().Foreground(lipgloss.Color("4")),
	TreeLinkName:        lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
	TreeMarkedNode: lipgloss.NewStyle().
		BorderLeft(true).
		BorderStyle(lipgloss.InnerHalfBlockBorder()).
		Background(lipgloss.Color("0")),
	TreeSelectionArrow: lipgloss.NewStyle().Foreground(lipgloss.Color("11")),
	TreeIndent:         lipgloss.NewStyle().Foreground(lipgloss.Color("0")),
	TreeIndentSelected: lipgloss.NewStyle().Foreground(lipgloss.Color("3")),

	PlainTextPreview: lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("15")),
}

var DefaultStylesheet = Stylesheet{
	SelectedPath: lipgloss.NewStyle().Foreground(lipgloss.Color("#74AC6D")),

	FinfoPermissions: lipgloss.NewStyle().Foreground(lipgloss.Color("#ACA46D")),
	FinfoLastUpdated: lipgloss.NewStyle().Foreground(lipgloss.Color("#E6E6E6")),
	FinfoSize:        lipgloss.NewStyle().Foreground(lipgloss.Color("#E6E6E6")),
	FinfoSep:         lipgloss.NewStyle().Foreground(lipgloss.Color("#2b2b2b")),

	OperationBar:      lipgloss.NewStyle().Foreground(lipgloss.Color("#E6E6E6")),
	OperationBarInput: lipgloss.NewStyle().Background(lipgloss.Color("#3C3C3C")),

	ErrBar:      lipgloss.NewStyle().Foreground(lipgloss.Color("#AC6D74")),
	HelpMsg:     lipgloss.NewStyle().Foreground(lipgloss.Color("#ACA46D")),
	HelpContent: lipgloss.NewStyle().Foreground(lipgloss.Color("#8c7ca6")),

	TreeRegularFileName: lipgloss.NewStyle().Foreground(lipgloss.Color("#E6E6E6")),
	TreeDirecotryName:   lipgloss.NewStyle().Foreground(lipgloss.Color("#6D74AC")),
	TreeLinkName:        lipgloss.NewStyle().Foreground(lipgloss.Color("#6DACA4")),
	TreeMarkedNode: lipgloss.NewStyle().
		BorderLeft(true).
		BorderStyle(lipgloss.InnerHalfBlockBorder()).
		Background(lipgloss.Color("#363636")),
	TreeSelectionArrow: lipgloss.NewStyle().Foreground(lipgloss.Color("#ACA46D")),
	TreeIndent:         lipgloss.NewStyle().Foreground(lipgloss.Color("#363636")),
	TreeIndentSelected: lipgloss.NewStyle().Foreground(lipgloss.Color("#ACA46D")),

	PlainTextPreview: lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#a8a8a8")),
}
