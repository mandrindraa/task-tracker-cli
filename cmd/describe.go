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
	fmt.Println("Task ID:", task.ID)
	fmt.Println("Task Name:", task.Name)
	fmt.Println("Task Status:", task.Status)
	fmt.Println("Task Created At:", task.CreatedAt.Format(time.ANSIC))
	fmt.Println("Task Updated At:", task.UpdatedAt.Format(time.ANSIC))
	fmt.Println(styles.SuccessIndication("Task details retrieved successfully"))
}
