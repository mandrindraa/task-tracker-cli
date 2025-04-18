package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Note: This example requires the lipgloss package to be installed:
// go get github.com/charmbracelet/lipgloss

// RunLipglossExample demonstrates how to use the lipgloss library for styling
// This function won't work until the lipgloss package is installed
func RunLipglossExample() {

	// Define styles

	taskStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		PaddingLeft(2)

	completedStyle := lipgloss.NewStyle().
		Strikethrough(true).
		Foreground(lipgloss.Color("#888888")).
		PaddingLeft(2)

	// Sample tasks
	tasks := []struct {
		ID        int
		Name      string
		Completed bool
	}{
		{1, "Implement task tracker CLI", false},
		{2, "Add beautiful console output", false},
		{3, "Write documentation", true},
	}

	// Render tasks
	for _, task := range tasks {
		taskID := fmt.Sprintf("%d.", task.ID)
		if task.Completed {
			fmt.Println(completedStyle.Render(taskID + " " + task.Name + " ✓"))
		} else {
			fmt.Println(taskStyle.Render(taskID + " " + task.Name))
		}
	}

}

func Title(content string) string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(2).
		PaddingRight(2).
		MarginBottom(1)
	return titleStyle.Render(content)
}

func SuccessIndication(content string) string {
	successTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#229900")).
		PaddingLeft(2).
		PaddingRight(2).
		MarginBottom(1)
	return successTitle.Render(content)
}
func ErrorIndication(content string) string {
	errorIndication := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#770000")).
		PaddingLeft(2).
		PaddingRight(2).
		MarginBottom(1)
	return errorIndication.Render(content)
}
