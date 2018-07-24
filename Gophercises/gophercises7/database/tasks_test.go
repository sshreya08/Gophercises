package database

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestSetupDB(t *testing.T) {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")
	err := SetupDB(dbPath)
	if err != nil {
		t.Errorf("Database path is incorrect, got: %d.", err)
	}
}

func TestSetupDBNegative(t *testing.T) {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks1.db")
	err := SetupDB(dbPath)
	if err != nil {
		t.Errorf("Database path is correct, got: %d.", err)
	}
}

func TestAddTask(t *testing.T) {
	id, err := AddTask("Laundary")
	if id == 0 {
		t.Errorf("Error in adding task: %d.", err)
	}
}

func TestShowTasks(t *testing.T) {
	var tasks []Task
	tasks, err := ShowTasks()
	if len(tasks) == 0 {
		t.Errorf("Error in viewing tasks: %d.", err)
	}
}

func TestDeleteTask(t *testing.T) {
	err := DeleteTask(1)
	if err != nil {
		t.Errorf("Error in viewing tasks: %d.", err)
	}
}
