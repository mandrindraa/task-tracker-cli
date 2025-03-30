package cmd

import (
	"fmt"
	"os"

	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
)

func init() {
	updateCmd := &cobra.Command{
		Use:   "update -i ID -n NEW_NAME -s NEW_STATUS",
		Short: "update infos about your task",
		Long:  "update the state or information about your task based on the task ID you provide",
		Run:   update,
	}
	updateCmd.Flags().IntP("id", "i", 1, "Update based on task ID")
	updateCmd.Flags().StringP("name", "n", "", "This will provide the NEW_NAME for the task")
	updateCmd.Flags().StringP("status", "s", "completed", "This will update task status")
	RootCmd.AddCommand(updateCmd)
}

func update(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		fmt.Println(styles.ErrorIndication("Invalid task ID"))
		os.Exit(1)
	}
	var task models.Task
	db.First(&task, "id = ?", id)
	status, err := cmd.Flags().GetString("status")
	if err != nil || status == "" || status == "\n" {
		fmt.Println(styles.ErrorIndication("Invalid operation: Task should have a status"))
		os.Exit(1)
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Println(styles.ErrorIndication("Invalid operation: Task should have a name"))
		os.Exit(1)
	}
	fmt.Print(name)

}
