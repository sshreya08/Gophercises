//package commands consists of functionality to implement command line through cobra in go
package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/sonam/Gophercises/gophercises7/database"
	"github.com/spf13/cobra"
)

//The list command will work on to List down all the tasks created
var doCommand = &cobra.Command{
	Use:   "done",
	Short: "Get All the done tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Please provide a task id to mark as complete.")
		}
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				log.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := database.ShowTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := database.DeleteTask(task.Serial_number)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

//Add the Done Command to Root Command
func init() {
	RootCommand.AddCommand(doCommand)
}
