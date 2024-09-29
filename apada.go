package main

import (
	"os",
	"fmt",
	"os/exec",
	"path/filepath"
	"strings",
	"text/templates"
)

const desktopTemplate = `[Desktop Entry]
TYPE=Application
Name={{.Name}}
Exec={{.Exec}}
Icon={{.Icon}}
Terminal=false
Categories={{Categories}}`

const DesktopEntry struct {
	Name		string
	Exec		string
	Icon		string
	Categories 	string
}

func main() {
	if len(os.Args) < 2 {
        fmt.Println("Usage: ./apada <filename>")
        os.Exit(1)
    }

	fileName = os.args[1]

	if *exec == "" {
		fmt.Printf("Error: Exec parameter is required")
	}
	file := readFile()
}

func readFile(name: string) {
	file, err := os.OpenFile("test.txt")

	if os.IsNotExist(err) {
		fmt.Printf("File not found")
		return
	}
	return file
}

func addExecutionPermission() {

}