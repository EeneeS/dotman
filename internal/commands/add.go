package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func Add(path string) {

	repoPath := os.ExpandEnv("$HOME/.dotfiles")

	if !isInitialized(repoPath) || !isPathExist(path) {
		return
	}

	absFilePath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	fileName := filepath.Base(absFilePath)
	targetPath := filepath.Join(repoPath, fileName)

	err = os.Rename(absFilePath, targetPath)
	if err != nil {
		fmt.Println("Error moving file/directory:", err)
		return
	}

	err = os.Symlink(targetPath, absFilePath)
	if err != nil {
		fmt.Println("Error creating symlink: ", err)
	}

	fmt.Println("Successfully added", fileName, "to dotman")
}

func isInitialized(repoPath string) bool {

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
