package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
)

func init() {
	listCmd := &cobra.Command{
		Use:   "list [-s Status]",
		Short: styles.Title("All task should appear whatever its status"),
		Long:  styles.Title("This will list all of your task"),
		Run:   list,
	}
	listCmd.Flags().StringP("status", "s", "", "Filter status of task")
	RootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	status, err := cmd.Flags().GetString("status")
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
	fmt.Println(styles.Bold + "╔═════╦════════════════════════════════╦═════════════╗" + styles.Reset)
	fmt.Printf(styles.Bold+"║ %-3s ║ %-30s ║ %-11s ║\n"+styles.Reset, "ID", "Task", "Status")
	fmt.Println(styles.Bold + "╠═════╬════════════════════════════════╬═════════════╣" + styles.Reset)
	if !rootVerbose {
		var tasks []models.Task
		if status != "" || status == "\n" {
			db.Select("id", "name", "status").Find(&tasks, "status = ?", strings.ToLower(status))
		} else {
			db.Select("id", "name", "status").Find(&tasks)
		}
		for _, task := range tasks {
			idStr := fmt.Sprintf("%d", task.ID)
			nameStr := task.Name
			if len(nameStr) > 30 {
				nameStr = nameStr[:27] + "..."
			}
			statusStr := task.Status
			if task.Status == models.Completed {
				nameStr = styles.Strikeout + styles.Green + nameStr + styles.Reset
				statusStr = styles.Green + statusStr + styles.Reset
			} else if task.Status == models.InProgress {
				statusStr = styles.Yellow + statusStr + styles.Reset
			} else if task.Status == models.Aborted {
				nameStr = styles.Red + nameStr + styles.Reset
				statusStr = styles.Red + statusStr + styles.Reset
			} else {
				statusStr = styles.Cyan + statusStr + styles.Reset
			}
			fmt.Printf("║ %-3s ║ %-30s ║ %-11s ║\n", idStr, nameStr, statusStr)
		}
	}

	fmt.Println(styles.Bold + "╚═════╩════════════════════════════════╩═════════════╝" + styles.Reset)
}
