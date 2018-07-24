//package cobra consists of functionality to implement command line through cobra in go
package cobra

import (
	"fmt"
	"log"

	"github.com/sonam/Gophercises/gophercises17"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{

	Use:   "set",
	Short: "Sets a secret in your secret storage. Please provide -k flag, followed by service_name, api key and key value space seperated.",
	Run: func(cmd *cobra.Command, args []string) {
		v := gophersises17.File(encodeKey, secretsPath())
		if len(args) == 0 || len(args) == 1 {
			log.Fatal("Insufficient arguments provided. Please try again...")
		}
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
