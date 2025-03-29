package examples

import (
	"fmt"
)

func RunExamples() {
	RunAnsiExample()
	RunLipglossExample()
	runTableExample()
}

func runTableExample() {
	// Define some ANSI color codes for the table
	const (
		reset     = "\033[0m"
		bold      = "\033[1m"
		cyan      = "\033[36m"
		green     = "\033[32m"
		yellow    = "\033[33m"
		gray      = "\033[90m"
		strikeout = "\033[9m"
	)

	// Sample tasks
	tasks := []struct {
		ID        int
		Name      string
		Status    string
		Completed bool
	}{
		{1, "Implement task tracker CLI", "In Progress", false},
		{2, "Add beautiful console output", "Pending", false},
		{3, "Write documentation", "Completed", true},
	}

	// Print table header
	fmt.Println(bold + "╔════╦═══════════════════════════════╦════════════╗" + reset)
	fmt.Printf(bold+"║ %s ║ %-30s ║ %-11s ║\n"+reset, "ID", "Task", "Status")
	fmt.Println(bold + "╠════╬═══════════════════════════════╬════════════╣" + reset)

	// Print table rows
	for _, task := range tasks {
		idStr := fmt.Sprintf("%d", task.ID)
		nameStr := task.Name
		statusStr := task.Status

		// Apply styling based on status
		if task.Completed {
			nameStr = strikeout + gray + nameStr + reset
			statusStr = green + statusStr + reset
		} else if task.Status == "In Progress" {
			statusStr = yellow + statusStr + reset
		} else {
			statusStr = cyan + statusStr + reset
		}

		fmt.Printf("║ %2s ║ %-30s ║ %-11s ║\n", idStr, nameStr, statusStr)
	}

	// Print table footer
	fmt.Println("╚════╩═══════════════════════════════╩════════════╝")
}
