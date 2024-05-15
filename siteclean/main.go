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

	fmt.Println(os.Args)
	// Get the directory specified as an argument
	rootDir := os.Args[1]
	targetDir := os.Args[2]

	// Get the absolute path of the target directory
	absTargetDir, err := filepath.Abs(targetDir)
	if err != nil {
		fmt.Printf("Error getting absolute path of %s: %v\n", targetDir, err)
		os.Exit(1)
	}
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		fmt.Printf("Error getting absolute path of %s: %v\n", rootDir, err)
		os.Exit(1)
	}

	//if true == true {
	//	os.Exit(0)
	//}
	// Walk through the directory
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Error getting absolute path of %s: %v\n", path, err)
			os.Exit(1)
		}

		fmt.Printf("path: %v \n absTargetDir: %v \n absRootDir: %v\n\n", absPath, absTargetDir, absRootDir)
		// Check if it's a directory (excluding the target directory)
		if info.IsDir() && absPath != absTargetDir && absPath != absRootDir {
			fmt.Printf("Deleting directory: %s\n", absPath)
			err := os.RemoveAll(absPath)
			if err != nil {
				fmt.Printf("Error deleting directory %s: %v\n", absPath, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Cleanup complete!")
}
