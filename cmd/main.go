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
	case "remove":
	case "help":
	default:
		// This will also be help
	}

}
