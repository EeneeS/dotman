package commands

import (
	"fmt"
	"os"
)

func Init() {
	repoPath := os.ExpandEnv("$HOME/.dotfiles")
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		err := os.Mkdir(repoPath, 0755)
		if err != nil {
			fmt.Println("Error creating repo:", err)
			return
		}
		fmt.Println("Initialized dotfile repository at:", repoPath)
	} else {
		fmt.Println("Repository already exists at:", repoPath)
	}
}
