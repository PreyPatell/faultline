package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// doctorModel is the Bubble Tea model for the interactive TUI.
type doctorModel struct {
	choices []string
	cursor  int
}

func newDoctorModel() doctorModel {
	return doctorModel{
		choices: []string{
			"List Services",
			"Inject Latency",
			"Degrade Network",
			"Kill Service",
			"Stop All Faults",
			"View Report",
			"Quit",
		},
	}
}

func (m doctorModel) Init() tea.Cmd { return nil }

func (m doctorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			if m.choices[m.cursor] == "Quit" {
				return m, tea.Quit
			}
			// TODO: wire up each selection to call the relevant chaos function
			fmt.Printf("\nSelected: %s (not yet wired up)\n", m.choices[m.cursor])
		}
	}
	return m, nil
}

func (m doctorModel) View() string {
	s := "=== Faultline Doctor ===\n"
	s += "Arrow keys to navigate • Enter to select • q to quit\n\n"
	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}
		s += fmt.Sprintf("%s%s\n", cursor, choice)
	}
	return s
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Interactive TUI for controlling chaos",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(newDoctorModel())
		_, err := p.Run()
		return err
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
