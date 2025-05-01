# RedisClone (Go)

A minimal Redis-like in-memory key-value store written in Go, supporting a subset of Redis commands and an Append-Only File (AOF) persistence mechanism.

## Features

- RESP protocol parser and serializer (compatible with Redis CLI)
- Supported commands:
  - `PING`
  - `SET key value`
  - `GET key`
  - `HSET hash key value`
  - `HGET hash key`
  - `HGETALL hash`
- Append-Only File (AOF) persistence:
  - All write operations (`SET`, `HSET`) are logged to disk
  - Replay on startup ensures data recovery

---

## Getting Started

### Prerequisites

- Go 1.18 or later installed

### Run the server

```bash
go run main.go
```

By default, the server listens on port 6379. You can connect using the Redis CLI:
```bash
redis-cli -p 6379
```

### Examples
```bash
127.0.0.1:6379> PING
"PONG"

127.0.0.1:6379> SET foo bar
"OK"

127.0.0.1:6379> GET foo
"bar"

127.0.0.1:6379> HSET user name Alice
"OK"

127.0.0.1:6379> HGET user name
"Alice"

127.0.0.1:6379> HGETALL user
1) "name"
2) "Alice"
```

### AOF Persistence
- The server writes every write command (SET, HSET) to an database.aof file.
- On startup, the file is parsed line-by-line and each command is re-executed in memory.
- This provides durability even after shutdown.

### How it works
Example AOF content:

```bash
*3
$3
SET
$3
foo
$3
bar
```

On server startup, this file is parsed and applied to restore the in-memory state.

### Code Structure
* `main.go`: entry point, server loop
* `handler.go`: command handlers (GET, SET, etc.)
* `resp.go`: RESP parser and writer
* `aof.go`: AOF write and replay logic