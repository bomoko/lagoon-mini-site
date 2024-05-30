package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if the argument is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <root> <directory>")
		os.Exit(1)
	}

	// Get the directory specified as an argument
	rootDir := os.Args[1]
	targetDir := os.Args[2]

	// Get the absolute path of the target directory
	absTargetDir := filepath.Join(rootDir, targetDir)

	_, err := os.Stat(absTargetDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Could not find safe directory '%v'\n Did you, maybe, specify an abosulte path for the target?\n", absTargetDir)
		} else {
			fmt.Print(err.Error())
		}
		os.Exit(1)
	}

	// Walk through the directory
	dirEntries, err := os.ReadDir(rootDir)

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	for _, e := range dirEntries {

		fullPath := filepath.Join(rootDir, e.Name())
		if e.IsDir() && fullPath != absTargetDir {
			fmt.Printf("Found dir to delete '%v' \n", fullPath)
			err := os.RemoveAll(fullPath)
			if err != nil {
				fmt.Printf("Unable to delete directory '%v' - failing\n", fullPath)
				os.Exit(1)
			}
		}

	}

	fmt.Println("Cleanup complete!")
}
