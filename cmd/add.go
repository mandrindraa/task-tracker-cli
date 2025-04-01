package cmd

import (
	"fmt"
	"os"

	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
)

func init() {
	addCmd := &cobra.Command{
		Use:   "add [-p 1] -n 'Wash the dishes' [-N 'Note']",
		Short: styles.Title("Add a new task"),
		Run:   add,
	}
	addCmd.Flags().StringP("name", "n", "", "Task name")
	addCmd.Flags().UintP("priority", "p", 100, "Task Priority")
	addCmd.Flags().StringP("note", "N", "", "Add A note to the Task")
	RootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if name == "" || err != nil {
		fmt.Println(styles.ErrorIndication("Error getting task name!"))
		os.Exit(1)
	}
	priority, err := cmd.Flags().GetUint("priority")
	if err != nil {
		fmt.Println(styles.ErrorIndication("Invalid Priority!"))
		os.Exit(1)
	}
	note, err := cmd.Flags().GetString("note")
	if err != nil {
		fmt.Println(styles.ErrorIndication("Invalid Note!"))
		os.Exit(1)
	}
	db.Create(&models.Task{
		Name:     name,
		Status:   models.ToDo,
		Note:     note,
		Priority: priority,
	})
	fmt.Println(styles.SuccessIndication("Adding task with name: " + name))

}
