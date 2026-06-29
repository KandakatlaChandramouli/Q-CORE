package main

import (
	"fmt"
	"os"

	"github.com/KandakatlaChandramouli/Q-CORE/internal/platform/version"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Q-CORE CLI")
		fmt.Println("Usage:")
		fmt.Println("  qcore-cli version")
		os.Exit(0)
	}

	switch args[0] {

	case "version":
		fmt.Println(version.Get().String())

	default:
		fmt.Printf("unknown command: %s\n", args[0])
		os.Exit(1)
	}
}
