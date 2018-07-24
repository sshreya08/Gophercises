//package database consists of functionality to implement add list and delete tasks in Bolt DB
package database

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

//Varible to create bucket for tasks
var taskBucketkey = []byte("tasks")

//Transaction parameter
var db *bolt.DB

//Create a key value struct
type Task struct {
	Serial_number int
	Task_name     string
}

// SetupDB Open the dbPath data file in the directory specified.
// It will be created if it doesn't exist.
func SetupDB(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucketkey)
		return err
	})
}

//AddTask adds tasks to the Task List which is a Bolt DB in this case and returns task id
func AddTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucketkey)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := inttobyte(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

//ShowTasks shows all pending tasks. It retrieves data from Bolt DB and returns a list of tasks
func ShowTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket(taskBucketkey)
		c := bk.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Serial_number: bytetoint(k),
				Task_name:     string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//DeleteTask marks a task as complete and removes it from Bolt DB based on tasks serial number
func DeleteTask(serial_number int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucketkey)
		return b.Delete(inttobyte(serial_number))
	})
}

//Convert int to byte
func inttobyte(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//Convert byte to int
func bytetoint(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
