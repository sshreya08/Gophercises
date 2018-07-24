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
2018/07/24 18:47:41 Task "Buy Groceries" added to the task list

$ Task add Do Laundary
2018/07/24 18:48:09 Task "Do Laundary" added to the task list.

$ Task list
Your Tasks:
1. Buy Groceries
2. Do Laundary

$ Task done 1
Marked "1" as completed.

$ Task list
You have the following tasks:
1. Do Laundary
```
