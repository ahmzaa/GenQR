package main

import (
    "os"
    "fmt"
//    "time"
//    "net/http"

    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var (
    // Available Spinners
    spinners = []spinner.Spinner{
        spinner.Line,
        spinner.Dot,
        spinner.MiniDot,
        spinner.Jump,
        spinner.Pulse,
        spinner.Points,
        spinner.Globe,
        spinner.Moon,
        spinner.Monkey,
    }

    textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Render
    spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
    helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)

type model struct {
    choices  []string
    cursor   int
    selected map[int]struct{}
    index    int
    spinner  spinner.Model
}

func initialModel() model {
    return model{
        // The to-do list, i.e list of tasks available to the user
        choices: []string{"Create QR Code"},

        // The map indicates which choices are selected.
        selected: make(map[int]struct{}),
    }
}

func (m model) Init() tea.Cmd {
    return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    // Is it a key press ?
    case tea.KeyMsg:
        switch msg.String() {
        // Keys to exit the program
        case "ctrl+c", "q":
            return m, tea.Quit
        // keys to move cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        // Keys to move cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        case "left", "h":
            m.index--
            if m.index < 0 {
                m.index = len(spinners) - 1
            }
            m.resetSpinner()
            return m, m.spinner.Tick
        case "right", "l":
            m.index++
            if m.index >= len(spinners) {
                m.index = 0
            }
            m.resetSpinner()
            return m, m.spinner.Tick
        // Keys to toggle the selected item / state
        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }

    // return the updated model to Bubble tea runtime
    // Note: no commands being returned.
    return m, nil
}

func (m *model) resetSpinner() {
    m.spinner = spinner.New()
    m.spinner.Style = spinnerStyle
    m.spinner.Spinner = spinners[m.index]
}

func (m model) View() string {
    // Header
    s := "QR Code Generator\n\n"

    // Iterate over choices
    for i, choice := range m.choices {
        
        // Is cursor pointing at this choice ?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor
        }

        // Is this choice selected?
        checked := " " // Note selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected
        }

        //render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    var gap string
    switch m.index {
    case 1:
        gap = ""
    default:
        gap = " "
    }

    s += fmt.Sprintf("\n %s%s%s\n\n", m.spinner.View(), gap, textStyle("Spinning..."))
    s += helpStyle("h/l, ←/→: change spinner • q: exit\n")

    // Footer
    s += "\nwritten by ahmza\n"
    s += "\nPress q to quit\n"

    return s
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("LoL, there's been an error: %v", err)
        os.Exit(1)
    }
}
