// Command durable-agents is the entry point for the runtime.
//
// Nothing works yet. See docs/ROADMAP.md for what comes first.
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
	fmt.Fprintln(os.Stderr, "durable-agents: nothing works yet. See docs/ROADMAP.md.")
	os.Exit(2)
}
