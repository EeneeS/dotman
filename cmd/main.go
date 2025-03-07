package main

import (
	"fmt"
	"os"

	"github.com/eenees/dotman/internal/commands"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: dotman <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		commands.Init()
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a file location")
			return
		}
		commands.Add(os.Args[2])
	case "remove":
	case "help":
	default:
		// This will also be help
	}

}
