# urn-lookup 🔍

urn-lookup is a web service for retrieving Redis hash data by URN key.

## Usage

Install the application using the `go` CLI or download the appropriate binary from the [releases page.](https://github.com/djaustin/urn-lookup/releases)

```bash
go install github.com/djaustin/urn-lookup@latest
```
## Configuration

The urn-lookup application is configured using environment variables. The following variables are accepted

|Variable|Required|Default|Description|
|-|:-:|-|-|
|PORT|No|8080|The port that urn-lookup will listen on|
|REDIS_ADDRESS|No|localhost:6379|The address of the Redis instance that will be searched|


