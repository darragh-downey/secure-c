package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/darragh-downey/secure-c/lexer"
	"github.com/darragh-downey/secure-c/parser"
	"github.com/darragh-downey/secure-c/semantic"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	highlightColor   = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle   = inactiveTabStyle.Copy().Bold(true)
	errorStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	warningStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))
	successStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	docStyle         = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	windowStyle      = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder())
)

type model struct {
	input        textinput.Model
	output       textarea.Model
	errMsg       string
	selectedFile string
	directory    filepicker.Model
	activeTab    int
	table        table.Model
	progress     progress.Model
	pager        paginator.Model
	help         help.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter C source code here"
	ti.Focus()
	ti.CharLimit = 500
	ti.Width = 50

	ta := textarea.New()
	ta.SetWidth(50)
	ta.SetHeight(20)
	// ta.SetStyle(windowStyle)

	directory := filepicker.New()
	directory.Path = "/home/ddowney/Workspace/github.com/secure_c/"

	tbl := table.New(table.WithColumns([]table.Column{
		{Title: "Line", Width: 5},
		{Title: "Issue", Width: 50},
	}))
	pgr := paginator.New()
	hp := help.New()
	prg := progress.New()

	return model{
		input:     ti,
		output:    ta,
		directory: directory,
		activeTab: 0,
		table:     tbl,
		progress:  prg,
		pager:     pgr,
		help:      hp,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.directory.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			if m.activeTab == 0 {
				output := processSourceCode(m.input.Value())
				m.output.SetValue(output)
				m.input.Reset()
			}
			return m, nil
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "tab":
			m.activeTab = (m.activeTab + 1) % 3
		case "left", "shift+tab":
			m.activeTab = (m.activeTab - 1 + 3) % 3
		}
	}

	switch m.activeTab {
	case 0:
		m.input, cmd = m.input.Update(msg)
	case 1:
		m.directory, cmd = m.directory.Update(msg)
	case 2:
		m.table, cmd = m.table.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	doc := strings.Builder{}

	var renderedTabs []string
	tabs := []string{"Editor", "Directory", "Issues"}
	for i, t := range tabs {
		var style lipgloss.Style
		if i == m.activeTab {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n\n")

	switch m.activeTab {
	case 0:
		doc.WriteString(inactiveTabStyle.Render("Enter your C code:") + "\n")
		doc.WriteString(m.input.View() + "\n\n")
		doc.WriteString(activeTabStyle.Render("Output:") + "\n")
		doc.WriteString(m.output.View())
	case 1:
		doc.WriteString(m.directory.View())
	case 2:
		doc.WriteString(m.table.View())
	}

	return docStyle.Render(doc.String())
}

func processSourceCode(source string) string {
	// Preprocess the source code
	preprocessedSource := lexer.Preprocess(source)

	// Lexical Analysis
	l := lexer.NewLexer(preprocessedSource)
	tokens := l.Tokenize()

	// Parsing
	p := parser.NewParser(tokens)
	ast, parseErr := p.Parse()
	if parseErr != nil {
		return errorStyle.Render(fmt.Sprintf("Parse error: %v", parseErr))
	}

	// Semantic Analysis
	semanticAnalyzer := semantic.NewAnalyzer()
	semanticErr := semanticAnalyzer.Analyze(ast)
	if semanticErr != nil {
		return errorStyle.Render(fmt.Sprintf("Semantic error: %v", semanticErr))
	}

	// Check for security issues
	securityIssues := checkForSecurityIssues(ast)
	if len(securityIssues) > 0 {
		return warningStyle.Render(fmt.Sprintf("Security issues:\n%s", formatSecurityIssues(securityIssues)))
	}

	return successStyle.Render("Code is valid and secure")
}

func checkForSecurityIssues(ast *parser.ASTNode) []string {
	var issues []string
	// Add logic to analyze the AST and identify security issues
	// For example:
	// - Unsafe functions like gets, strcpy, system, etc.
	// - Buffer overflows
	// - Format string vulnerabilities
	// - Integer overflows
	// - Unsafe memory management

	// Example check
	if ast.Value == "function_declaration" {
		for _, child := range ast.Children {
			if child.Value == "gets" || child.Value == "strcpy" || child.Value == "system" {
				issues = append(issues, fmt.Sprintf("Unsafe function usage: %s", child.Value))
			}
		}
	}

	return issues
}

func formatSecurityIssues(issues []string) string {
	var formatted strings.Builder
	for _, issue := range issues {
		formatted.WriteString(fmt.Sprintf("- %s\n", issue))
	}
	return formatted.String()
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
