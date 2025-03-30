package styles

import (
	"fmt"
)

func RunExamples() {
	RunLipglossExample()
	runTableExample()
}

func runTableExample() {
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
	fmt.Println(Bold + "╔════╦═══════════════════════════════╦════════════╗" + Reset)
	fmt.Printf(Bold+"║ %s ║ %-30s ║ %-11s ║\n"+Reset, "ID", "Task", "Status")
	fmt.Println(Bold + "╠════╬═══════════════════════════════╬════════════╣" + Reset)

	// Print table rows
	for _, task := range tasks {
		idStr := fmt.Sprintf("%d", task.ID)
		nameStr := task.Name
		statusStr := task.Status

		// Apply styling based on status
		if task.Completed {
			nameStr = Strikeout + Green + nameStr + Reset
			statusStr = Green + statusStr + Reset
		} else if task.Status == "In Progress" {
			statusStr = Yellow + statusStr + Reset
		} else {
			statusStr = Cyan + statusStr + Reset
		}

		fmt.Printf("║ %2s ║ %-30s ║ %-11s ║\n", idStr, nameStr, statusStr)
	}

	// Print table footer
	fmt.Println("╚════╩═══════════════════════════════╩════════════╝")
}
