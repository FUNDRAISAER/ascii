package main

import (
    "fmt"
    "os"
    "strings"
)

var colors = map[string]string{
    "red":    "\033[31m",
    "green":  "\033[32m",
    "yellow": "\033[33m",
    "blue":   "\033[34m",
    "reset":  "\033[0m",
}

func usage() {
    fmt.Println("Usage: go run . [OPTION] [STRING]")
    fmt.Println()
    fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
}

func colorString(colorName, text string) (string, error) {
    colorCode, ok := colors[colorName]
    if !ok {
        return "", fmt.Errorf("invalid color: %s", colorName)
    }
    return colorCode + text + colors["reset"], nil
}

func colorSubstring(colorName, substring, text string) (string, error) {
    colorCode, ok := colors[colorName]
    if !ok {
        return "", fmt.Errorf("invalid color: %s", colorName)
    }
    return strings.ReplaceAll(
        text,
        substring,
        colorCode+substring+colors["reset"],
    ), nil
}

func main() {
    args := os.Args

    if len(args) < 3 || len(args) > 4 {
        usage()
        return
    }

    if !strings.HasPrefix(args[1], "--color=") {
        usage()
        return
    }

    colorName := strings.Split(args[1], "=")[1]

    if len(args) == 3 {
        result, err := colorString(colorName, args[2])
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println(result)
        return
    }

    result, err := colorSubstring(colorName, args[2], args[3])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
}