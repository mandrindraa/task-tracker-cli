package cmd

import (
	"fmt"
	"strings"

	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
)

func init() {
	updateCmd := &cobra.Command{
		Use:   "update -i ID -n NEW_NAME -s NEW_STATUS",
		Short: styles.Title("update infos about your task"),
		Long:  "update the state or information about your task based on the task ID you provide",
		Run:   update,
	}
	updateCmd.Flags().IntP("id", "i", 1, "Update based on task ID")
	updateCmd.Flags().StringP("name", "n", "", "This will provide the NEW_NAME for the task")
	updateCmd.Flags().StringP("status", "s", "completed", "This will update task status")
	updateCmd.Flags().String("no-status-change", "", "This will not change the status of the task")
	RootCmd.AddCommand(updateCmd)
}

func update(cmd *cobra.Command, args []string) {
	id, err := getTaskID(cmd)
	if err != nil {
		HandleError("Invalid task ID", err)
	}

	task, err := findTaskByID(id)
	if err != nil {
		HandleError(err.Error(), nil)
	}

	status, err := getStatus(cmd)
	if err != nil {
		HandleError("Invalid status", err)
	}

	name, err := getName(cmd)
	if err != nil {
		HandleError("Invalid name", err)
	}

	if err := updateTaskFields(&task, name, status); err != nil {
		HandleError("Failed to update the task", err)
	}

	fmt.Println(styles.SuccessIndication("Task updated successfully"))
}

// getTaskID retrieves the task ID from the command flags
func getTaskID(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("id")
}

// findTaskByID retrieves a task by its ID from the database

// getStatus retrieves and validates the status from the command flags
func getStatus(cmd *cobra.Command) (string, error) {
	status, err := cmd.Flags().GetString("status")
	if err != nil {
		return "", fmt.Errorf("task should have a valid status")
	}
	validStatus, ok := models.IsValidStatus(status)
	if !ok {
		return "", fmt.Errorf("invalid status! Allowed values are: pending, in progress, completed, aborted")
	}
	return strings.ToLower(validStatus), nil
}

// getName retrieves the name from the command flags
func getName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("name")
}

// updateTaskFields updates the task fields and saves it to the database
func updateTaskFields(task *models.Task, name, status string) error {
	if name != "" {
		task.Name = name
	}
	task.Status = status

	result := db.Save(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
