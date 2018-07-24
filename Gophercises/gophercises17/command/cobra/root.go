//package cobra consists of functionality to implement command line through cobra in go
package cobra

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

//Runs root command
var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is an API key and other secrets manager",
}

var encodeKey string

//Use flag for secret key through command line
func init() {
	RootCmd.PersistentFlags().StringVarP(&encodeKey, "key", "k", "", "the key to use when encoding and decoding secrets")
}

//Set path for secret file in Home Directory
func secretsPath() string {
	//home, _ := homedir.Dir()
	return filepath.Join(".secrets")
}
