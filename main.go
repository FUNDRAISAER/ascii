package main

import (
    "fmt"
    "os"
    "strings"
)
// this is a map that contains the color names as keys and their corresponding ANSI escape codes as values
var colors = map[string]string{
    "red":    "\033[31m",
    "green":  "\033[32m",
    "yellow": "\033[33m",
    "blue":   "\033[34m",
    "reset":  "\033[0m",
}

// this function prints the usage instuctions for the program, it is called when the user provides invaild arguments, so the user can understand how to use the program correctly.
func usage() {
    fmt.Println("Usage: go run . [OPTION] [STRING]")
    fmt.Println()
    fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
}

// this function takes a color name and a string as input, it checks if the color name is valid by looking it up in the colors map, if it is valid it returns the string wrapped in the corresponding ANSI espae code to color the entire string, if the color name is invalid it returns an error.
func colorString(colorName, text string) (string, error) {
    colorCode, ok := colors[colorName]
    if !ok {
        return "", fmt.Errorf("invalid color: %s", colorName)
    }
    // this returns the text wrapped in the color code and reset code, so that the text will be colored when printed to the terminal, and after the text is printed the color will be reset to default.
    return colorCode + text + colors["reset"], nil
}
// this function takes a color name, a substring, and a text as input, it checks if the color name is valid by looking it up in the colors map, if it is valid it replaces all occurrences of the substring in the text with the substring wrapped in the corresponding ANSI escape code to color only the substring, if the color name is invalid it returns an error.
func colorSubstring(colorName, substring, text string) (string, error) {
    colorCode, ok := colors[colorName]
    if !ok {
        return "", fmt.Errorf("invalid color: %s", colorName)
    }
    // this replaces all occurrences of the substring in the text with the substring wrapped in the color code and reset code, so that only the substring will be colored when printed to the terminal, and after the substring is printed the color will be reset to default.
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
    // this checks if the second argument starts with "--color="
    if !strings.HasPrefix(args[1], "--color=") {
        usage()
        return
    }
    // this extracts or pick the color name from the second argument, by picking the part after the "=" sign and store it in the variable colorName
    // for example the --color is [0] and the color name is [1] so now the [1] is the color name
    colorName := strings.Split(args[1], "=")[1]
// this checks if the number of arguments is 3, if it is 3 it means that the user wants to color the entire string, so it calls the colorString function with the color name and the string to be colored, if there is an error it prints the error and retturns, if there is no error it prints the result and returns.
    if len(args) == 3 {
        result, err := colorString(colorName, args[2])
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println(result)
        return
    }
// if the number of arguments is 4, it means that the user wants to color only a substring of the string, so it calls the colorSubstring function with the color name, the substring to be colored, and the string to be colored, if there is an error it prints the error and returns, if there is no error it prints the result and returns.
    result, err := colorSubstring(colorName, args[2], args[3])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
}
