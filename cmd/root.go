package cmd

import (
	"fmt"

	"github.com/mandrindraa/task-tracker-cli/database"
	"github.com/mandrindraa/task-tracker-cli/models"
	"github.com/mandrindraa/task-tracker-cli/styles"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var rootVerbose bool
var db *gorm.DB

var RootCmd = &cobra.Command{
	Use:   "tt-cli",
	Short: "To Do list manager in your Terminal! Use --help for more information",
	Long:  "Task-Tracker-Cli is a command-line interface application for managing your to-do list in a terminal user interface.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.Title("To Do list manager in your Terminal! Use --help for more information"))
	},
}

func init() {
	db = database.GetDB()
	db.AutoMigrate(&models.Task{})
	RootCmd.PersistentFlags().BoolVarP(&rootVerbose, "verbose", "v", false, "enable verbose output")
}
