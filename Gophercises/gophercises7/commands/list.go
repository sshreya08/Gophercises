//package commands consists of functionality to implement command line through cobra in go
package commands

import (
	"fmt"
	"log"

	"github.com/sonam/Gophercises/gophercises7/database"
	"github.com/spf13/cobra"
)

//The list command will work on to List down all the tasks created
var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List All the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := database.ShowTasks()
		if err != nil {
			log.Fatal("Could not List", err.Error())
		}
		if len(tasks) == 0 {
			log.Println("No pending tasks!!")
			return
		}
		fmt.Println("Your Tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Task_name)
		}
	},
}

//Add the List Command to Root Command
func init() {
	RootCommand.AddCommand(listCommand)
}
