package main

import (
	"runtime"

	"github.com/nullsploit01/cc-redis/server/cmd"
)

func main() {
	runtime.GOMAXPROCS(5) // allocates 5 threads
	cmd.Execute()
}
