# TASK MANAGER
In this exercise we are creating command line tool for todo tasks.

```
We can perform following tasks through this
1. Add task to the tasks list
2. List tasks from the tasks list
3. Mark a task as complete
```

```
Commands:
Task
This Application is a Command Line Task Manager

Usage:
  Task [command]

Available Commands:
  add         Add All the tasks
  done        Get All the done tasks
  help        Help about any command
  list        List All the tasks

Flags:
  -h, --help   help for Task

Use "Task [command] --help" for more information about a commands

$ Task add Buy Groceries
Added "review talk proposal"

$ Task add Do Laundary
Added "clean dishes"

$ Task list
1. review talk proposal
2. some task description

$ Task done 1

$ Task list
You have the following tasks:
1. some task description
```
