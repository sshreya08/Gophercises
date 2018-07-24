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
Secret is an API key and other secrets manager

Usage:
  secret [command]

Available Commands:
  get         Gets a secret in your secret storage. Please provide -k flag, followed by service_name and api key
  help        Help about any command
  set         Sets a secret in your secret storage. Please provide -k flag, followed by service_name, api key and key value space seperated.

Flags:
  -h, --help         help for secret
  -k, --key string   the key to use when encoding and decoding secrets

Use "secret [command] --help" for more information about a command.

$ secret set -k api-key key-name key-value
Value set successfully!

$ secret get -k api-key key-name
key-name = key-value
```
