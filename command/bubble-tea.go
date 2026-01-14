package command

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// 菜单选项
type menuItem struct {
	title       string
	description string
}

// Model 定义菜单的状态
type model struct {
	cursor   int
	selected int
	choices  []string
}

func initialModel() model {
	return model{
		cursor:   0,
		selected: -1,
		choices: []string{
			"添加新密码",
			"查看密码列表",
			"搜索密码",
			"删除密码",
			"退出",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	// The header
	s := "极光密码管理工具, 请选择操作:\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if m.selected == i {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "enter":
			m.selected = m.cursor
		}
	}
	return m, nil
}

func ExcuteBubbleTeaMenu() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
