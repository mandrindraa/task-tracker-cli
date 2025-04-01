package cmd

import (
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
)

func init() {
	removeCmd := &cobra.Command{
		Use:   "remove {-i <ID> | -s <Status>}",
		Short: styles.Title("Remove task based on id or status"),
		Long:  "Remove task if it is no longer relevant",
		Run:   remove,
	}
	removeCmd.Flags().UintP("id", "i", 1, "provide task id to remove")
	removeCmd.Flags().StringP("status", "s", "completed", "provide the task status to be removed")
	RootCmd.AddCommand(removeCmd)
}

func remove(cmd *cobra.Command, args []string) {
}
