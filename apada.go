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

    fileName := os.Args[1]
    file, err := readFile(fileName)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
}

func readFile(name string) (*os.File, error) {
    file, err := os.Open(name)

    if err != nil {
        if os.IsNotExist(err) {
            return nil, fmt.Errorf("file not found")
        }
        return nil, err
    }
    return file, nil
}

func addExecutionPermission() {

}
