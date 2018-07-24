//package main
package main

import (
	"log"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sonam/Gophercises/gophercises7/commands"
	"github.com/sonam/Gophercises/gophercises7/database"
)

//Invoking Cobra Tasks from here
func main() {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")
	check(database.SetupDB(dbPath))
	check(commands.RootCommand.Execute())
}

//Put a check if any error shown
func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
