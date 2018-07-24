//package cobra
//package cobra consists of functionality to implement command line through cobra in go
package cobra

import (
	"fmt"
	"log"

	"github.com/sonam/Gophercises/gophercises17"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your secret storage. Please provide -k flag, followed by service_name and api key",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("key=%s\n", encodingKey)
		v := gophersises17.File(encodeKey, secretsPath())
		if len(args) == 0 {
			log.Fatal("Insufficient arguments provided. Please try again...")
		}
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
