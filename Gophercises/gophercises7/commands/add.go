//package commands consists of functionality to implement command line through cobra in go
package commands

import (
	"log"
	"strings"

	"github.com/sonam/Gophercises/gophercises7/database"

	"github.com/spf13/cobra"
)

//The list command will work on to List down all the tasks created
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add All the tasks",
	Long:  "Add tasks to the task bucket, one at a time",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		if len(args) == 0 {
			log.Fatal("Please provide a task to add to the list")
		}
		_, err := database.AddTask(task)
		if err != nil {
			log.Fatal("Could not add the task:", err)
		}
		log.Printf("Task \"%s\" added to the task list.\n", task)

	},
}

//Add the Add Command to Root Command
func init() {
	RootCommand.AddCommand(addCommand)
}
