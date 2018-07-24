# SECRET API KEY MANAGER
In this exercise we are managing Secret API key and other secrets

```
We can perform following tasks through this
1. Set API Key
2. Get API Key
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
