package command

import tea "github.com/charmbracelet/bubbletea"

// 菜单选项
type menuItem struct {
	title       string
	description string
}

var menuItems = []menuItem{
	{"添加新密码", "添加一个新的密码条目"},
	{"查看密码列表", "查看所有已保存的密码"},
	{"搜索密码", "根据关键词搜索密码"},
	{"删除密码", "删除指定的密码条目"},
	{"退出", "退出程序"},
}

// Model 定义菜单的状态
type model struct {
	cursor   int
	selected int
	items    []menuItem
}

func (m model) Init() tea.Cmd {
	return nil
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
				m.cursor = len(m.items) - 1
			}
		case "down":
			if m.cursor < len(m.items)-1 {
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
