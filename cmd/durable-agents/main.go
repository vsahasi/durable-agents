// Command durable-agents is the entry point for the runtime.
package main

import (
	"fmt"
	"os"
)

const version = "0.0.0"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("durable-agents", version)
		return
	}
	fmt.Fprintln(os.Stderr, "durable-agents: nothing runs yet")
	os.Exit(2)
}
