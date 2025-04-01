package cmd

import (
	"fmt"
	"time"

	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
)

func init() {
	describeCmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe the task tracker CLI",
		Long:  "This command provides a detailed description of the task tracker CLI.",
		Run:   describe,
	}
	describeCmd.Flags().IntP("id", "i", 0, "Provide the task ID to describe")
	describeCmd.MarkFlagRequired("id")
	RootCmd.AddCommand(describeCmd)
}

func describe(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		HandleError("Task ID is required", err)
	}
	task, err := findTaskByID(id)
	if err != nil {
		HandleError("Task not found", err)
	} else {
		printTaskDetails(task)
	}
}

func printTaskDetails(task models.Task) {
	fmt.Println(styles.Bold + "╔═════════════════════════════════════════════════════════╗" + styles.Reset)
	fmt.Println(styles.Bold + styles.Cyan + "║                     Task Details                        ║" + styles.Reset)
	fmt.Println(styles.Bold + "╠════════════════════════╦════════════════════════════════╣" + styles.Reset)
	fmt.Printf(styles.Bold+"║ %-22s ║ %-30s ║\n"+styles.Reset, "Field", "Value")
	fmt.Println(styles.Bold + "╠════════════════════════╬════════════════════════════════╣" + styles.Reset)

	// Print task details with wrapping for long values
	printField("Task ID", fmt.Sprintf("%d", task.ID))
	printField("Task Name", task.Name)
	printField("Task Status", task.Status)
	printField("Task Note", task.Note)
	printField("Task Created At", task.CreatedAt.Format(time.ANSIC))
	printField("Task Updated At", task.UpdatedAt.Format(time.ANSIC))

	fmt.Println(styles.Bold + "╚════════════════════════╩════════════════════════════════╝" + styles.Reset)
}

// printField prints a field and its value, wrapping the value if it is too long
func printField(field, value string) {
	const maxWidth = 30 // Maximum width for a single line
	if len(value) > maxWidth {
		// Split the value into multiple lines
		lines := splitIntoLines(value, maxWidth)
		fmt.Printf("║ %-22s ║ %-30s ║\n", field, lines[0])
		for _, line := range lines[1:] {
			fmt.Printf("║ %-22s ║ %-30s ║\n", "", line)
		}
	} else {
		fmt.Printf("║ %-22s ║ %-30s ║\n", field, value)
	}
}

// splitIntoLines splits a string into multiple lines of a given maximum width
func splitIntoLines(value string, maxWidth int) []string {
	var lines []string
	for len(value) > maxWidth {
		lines = append(lines, value[:maxWidth])
		value = value[maxWidth:]
	}
	lines = append(lines, value)
	return lines
}
