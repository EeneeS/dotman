package commands

import (
	"fmt"
	"os"
)

func Add(path string) {

	if !isInitialized() {
		return
	}

	if !isPathExist(path) {
		return
	}

	fmt.Println(path)
}

func isInitialized() bool {
	repoPath := os.ExpandEnv("$HOME/.dotfiles")

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		fmt.Println("Error: Dotfile repository does not exist at", repoPath)
		fmt.Println("Please run 'dotman init' to initialize the repository.")
		fmt.Println("For more information, run 'dotman help'.")
		return false
	}

	return true
}

func isPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Path not found:", path)
		return false
	}

	return true
}
