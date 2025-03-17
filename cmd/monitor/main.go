// Declare Package
package main

// Make Imports
import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Implement Stand In Model
type model struct {
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return "Hello World"
}

// Main Function
func main() {

	// Create New Program
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Errorf("There has been an error setting up the app: %v", err)
	}
}
