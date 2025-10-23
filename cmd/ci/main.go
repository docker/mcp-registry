package main

import (
	"fmt"
	"os"
)

// main dispatches the CLI to a specific sub-command implementation.
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: ci <command> [options]")
		os.Exit(2)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var err error
	switch cmd {
	case "collect-updated-pins":
		err = runCollectUpdatedPins(args)
	case "prepare-updated-pins":
		err = runPrepareUpdatedPins(args)
	case "collect-new-servers":
		err = runCollectNewServers(args)
	case "prepare-new-servers":
		err = runPrepareNewServers(args)
	case "compose-pr-summary":
		err = runComposePRSummary(args)
	case "collect-full-audit":
		err = runCollectFullAudit(args)
	case "prepare-full-audit":
		err = runPrepareFullAudit(args)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
