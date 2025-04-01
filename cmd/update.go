package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
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
	RootCmd.AddCommand(updateCmd)
}

func update(cmd *cobra.Command, args []string) {
	id, err := getTaskID(cmd)
	if err != nil {
		handleError("Invalid task ID", err)
	}

	task, err := findTaskByID(id)
	if err != nil {
		handleError(err.Error(), nil)
	}

	status, err := getStatus(cmd)
	if err != nil {
		handleError("Invalid status", err)
	}

	name, err := getName(cmd)
	if err != nil {
		handleError("Invalid name", err)
	}

	if err := updateTaskFields(&task, name, status); err != nil {
		handleError("Failed to update the task", err)
	}

	fmt.Println(styles.SuccessIndication("Task updated successfully"))
}

// getTaskID retrieves the task ID from the command flags
func getTaskID(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("id")
}

// findTaskByID retrieves a task by its ID from the database
func findTaskByID(id int) (models.Task, error) {
	var task models.Task
	result := db.First(&task, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return task, fmt.Errorf("task not found")
		}
		return task, fmt.Errorf("error retrieving task: %v", result.Error)
	}
	return task, nil
}

// getStatus retrieves and validates the status from the command flags
func getStatus(cmd *cobra.Command) (string, error) {
	status, err := cmd.Flags().GetString("status")
	if err != nil || status == "" || status == "\n" {
		return "", fmt.Errorf("task should have a valid status")
	}
	if !models.IsValidStatus(status) {
		return "", fmt.Errorf("invalid status! Allowed values are: pending, in progress, completed, aborted")
	}
	return strings.ToLower(status), nil
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

// handleError prints an error message and exits the program
func handleError(message string, err error) {
	if err != nil {
		fmt.Println(styles.ErrorIndication(fmt.Sprintf("%s: %v", message, err)))
	} else {
		fmt.Println(styles.ErrorIndication(message))
	}
	os.Exit(1)
}
