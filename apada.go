package main

import (
    "fmt"
    "os"
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
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./apada <filename>")
        os.Exit(1)
    }

    fileName := os.Args[1] // Declare fileName properly
    file, err := readFile(fileName)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close() // Don't forget to close the file when you're done
}

func readFile(name string) (*os.File, error) { // Correct parameter syntax
    file, err := os.Open(name) // Open the file

    if err != nil {
        if os.IsNotExist(err) {
            return nil, fmt.Errorf("file not found") // Return error with a message
        }
        return nil, err // Return the error for other cases
    }
    return file, nil // Return the file and nil error
}

func addExecutionPermission() {
    // Implementation goes here
}
