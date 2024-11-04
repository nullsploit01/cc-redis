# CC-Redis: Custom Redis Server CLI

CC-Redis is a command-line interface (CLI) designed to manage a custom Redis server. This project is part of a coding challenge from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-redis/).

## Project Overview

CC-Redis provides tools to launch and interact with a Redis server, supporting custom configurations to suit various deployment environments. The CLI is designed for efficiency and flexibility, allowing users to specify server options directly from the command line.

## Prerequisites

Before running CC-Redis, ensure you have the following installed:

- **Redis**: CC-Redis requires Redis installed on your machine for full functionality. Download and install it from [Redis's official website](https://redis.io/download).
- **Go (at least version 1.13)**: CC-Redis is written in Go. Download and install it from [Go's official website](https://golang.org/dl/).

## Features

- Interactive CLI mode for executing Redis commands like `SET`, `GET`, `DEL`, `PING`, and `ECHO`.
- Communication via **RESP (REdis Serialization Protocol)**, making it compatible with Redis clients.
- Configurable options for server port and host.
- Benchmark testing for server performance.

## Installation

To install CC-Redis, follow these steps:

```bash
git clone https://github.com/nullsploit01/cc-redis.git
cd cc-redis
go build -o ccredis-server ./server
go build -o ccredis-cli ./cli
```

## Usage

To start the Redis server, run:

```bash
./ccredis-server --port 6379 --host localhost
```

To connect to the server and enter interactive mode, use:

```bash
./ccredis-cli --port 6379 --host localhost
```

Once connected, you will see a prompt to start issuing commands, similar to the following:

```bash
Connected to ccredis-server at localhost:6379
> SET key value
OK
> GET key
value
> DEL key
OK
> PING
PONG
> ECHO "Hello, CC-Redis!"
Hello, CC-Redis!
```

## Benchmark Testing

To test the performance of SET and GET commands, use the redis-benchmark tool. This command will run 100,000 operations for SET and GET and display the results:

```bash
redis-benchmark -t set,get -n 100000 -q
```

## Benchmark Results

```bash
SET: 206611.58 requests per second, p50=0.127 msec
GET: 206185.56 requests per second, p50=0.127 msec
```

These results indicate that CC-Redis achieves high throughput for basic operations, with minimal latency under typical workloads.
