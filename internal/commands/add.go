package commands

import (
	"fmt"
	"os"
)

func Add(path string) {

	repoPath := os.ExpandEnv("$HOME/.dotfiles")

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		fmt.Println("Error: Dotfile repository does not exist at", repoPath)
		fmt.Println("Please run 'dotman init' to initialize the repository.")
		fmt.Println("For more information, run 'dotman help'.")
		return
	}

	fmt.Println(path)
}
