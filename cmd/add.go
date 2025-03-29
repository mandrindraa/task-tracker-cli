package cmd

import (
	"fmt"
	"os"

	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/spf13/cobra"
)

func init() {
	addCmd := &cobra.Command{
		Use:   "add -t 1 -n 'Wash the dishes'",
		Short: "Add a task",
		Run:   add,
	}
	addCmd.Flags().StringP("name", "n", "", "Task name")
	RootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	if name, err := cmd.Flags().GetString("name"); err != nil || name == "" {
		fmt.Println("Error getting task name:", err)
		os.Exit(1)
	} else {
		db.Create(&models.Task{
			Name:   name,
			Status: models.Status("To Do"),
			Note:   "",
		})
		fmt.Println("Adding task with name:", name)
	}
}
