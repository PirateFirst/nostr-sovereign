package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("sp — Sovereign Pathway CLI")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  sp ingest <file.md>")
		fmt.Println("  sp materialize <object-id>")
		fmt.Println("")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "ingest":
		fmt.Println("ingest: not yet implemented")
	case "materialize":
		fmt.Println("materialize: not yet implemented")
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
