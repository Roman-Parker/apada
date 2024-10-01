package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const desktopTemplate = `[Desktop Entry]
TYPE=Application
Name={{.Name}}
Exec={{.Exec}}
Icon={{.Icon}}
Terminal=false
Categories={{.Categories}}`

type DesktopEntry struct {
	Name       string
	Exec       string
	Icon       string
	Categories string
}

func main() {
	// Check if the fileName argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./apada <filename>")
		os.Exit(1)
	}

	// Read command-line flags
	sourceFile := flag.String("file", "", "Path to the source file")
	flag.Parse()

	fileName := os.Args[1]
	targetDir := "/home/Applications/AppImages"

	moveFileToTargetDirectory(targetDir, *sourceFile, fileName)

	addExecutionPermission(targetFile)
}

// moveFileToTargetDirectory handles moving the file to the target directory
func moveFileToTargetDirectory(targetDir string, sourceFile string, fileName string) {
	// If no source file is provided, assume it's in the current directory
	if sourceFile == "" {
		currentDir, err := os.Getwd() // Get current working directory
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			os.Exit(1)
		}

		// Set the source file path to the current directory + fileName
		sourceFile = filepath.Join(currentDir, fileName)
	}

	// Check if the target directory exists, create it if not
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		// Create the target directory if it doesn't exist
		err := os.MkdirAll(targetDir, 0755)
		if err != nil {
			fmt.Println("Error creating target directory:", err)
			os.Exit(1)
		}
		fmt.Println("Target directory created:", targetDir)
	}

	// Define the target file path
	targetFile := filepath.Join(targetDir, filepath.Base(sourceFile))

	// Move the file to the target directory
	err := os.Rename(sourceFile, targetFile)
	if err != nil {
		fmt.Println("Error moving file:", err)
		os.Exit(1)
	}

	fmt.Println("File successfully moved to:", targetFile)
    return targetFile
}

// addExecutionPermission adds execution permissions to the file
func addExecutionPermission(filePath string) {
	err := os.Chmod(filePath, 0755)
	if err != nil {
		fmt.Println("Error adding execution permissions to the file:", err)
		os.Exit(1)
	}
	fmt.Println("Execution permissions added to:", filePath)
}
