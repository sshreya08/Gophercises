//package commands consists of functionality to implement command line through cobra in go
package commands

import (
	"github.com/spf13/cobra"
)

//The root command represents a bare call to the Task application.
var RootCommand = &cobra.Command{
	Use:   "Task",
	Short: "This Application is a Command Line Task Manager",
}
