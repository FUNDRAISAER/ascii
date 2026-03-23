package main

import (
    "testing"
)
// these are unit tests for the colorString and colorSubstring functions, they test different scenarios including valid colors, invalid colors, and edge cases to ensure that the functions behave as expected in various situations.
func TestColorString(t *testing.T) {
    tests := []struct {
        name      string
        color     string
        text      string
        expected  string
        shouldErr bool
    }{
        {
            name:     "Red color",
            color:    "red",
            text:     "Hello",
            expected: "\033[31mHello\033[0m",
        },
        {
            name:     "Green color",
            color:    "green",
            text:     "World",
            expected: "\033[32mWorld\033[0m",
        },
        {
            name:     "Yellow color",
            color:    "yellow",
            text:     "Test",
            expected: "\033[33mTest\033[0m",
        },
        {
            name:     "Blue color",
            color:    "blue",
            text:     "ASCII",
            expected: "\033[34mASCII\033[0m",
        },
        {
            name:      "Invalid color",
            color:     "purple",
            text:      "Hello",
            shouldErr: true,
        },
    }
// this loop iterates over the test cases defined in the tests slice, for each test case it runs a subtest using t.Run with the name of the test case, it calls the colorString function with the color and text from the test case, it checks if an error was returned when it was expected or not, and if there was no error it checks if the result matches the expected output, if any of these checks fail it reports an error using t.Errorf.
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := colorString(tt.color, tt.text)
            if (err != nil) != tt.shouldErr {
                t.Errorf("colorString() error = %v, shouldErr %v", err, tt.shouldErr)
            }
            if !tt.shouldErr && result != tt.expected {
                t.Errorf("colorString() = %q, want %q", result, tt.expected)
            }
        })
    }
}
// this test function tests the colorSubstring function with various scenarios, it checks if the function correctly colors the specified substring in the text, handles multiple occurrences of the substring, and returns an error for invalid colors. It uses a similar structure to the TestColorString function, iterating over a set of test cases and checking the results against expected values.
func TestColorSubstring(t *testing.T) {
    tests := []struct {
        name       string
        color      string
        substring  string
        text       string
        expected   string
        shouldErr  bool
    }{
        {
            name:      "Color substring in text",
            color:     "red",
            substring: "lo",
            text:      "Hello",
            expected:  "Hel\033[31mlo\033[0m",
        },
        {
            name:      "Color multiple occurrences",
            color:     "green",
            substring: "a",
            text:      "banana",
            expected:  "b\033[32ma\033[0mn\033[32ma\033[0mn\033[32ma\033[0m",
        },
        {
            name:      "Substring not found",
            color:     "blue",
            substring: "xyz",
            text:      "Hello",
            expected:  "Hello",
        },
        {
            name:      "Invalid color",
            color:     "orange",
            substring: "test",
            text:      "test",
            shouldErr: true,
        },
    }
// this loop iterates over the test cases defined in the tests slice, for each test case it runs a subtest using t.Run with the name of the test case, it calls the colorSubstring function with the color, substring, and text from the test case, it checks if an error was returned when it was expected or not, and if there was no error it checks if the result matches the expected output, if any of these checks fail it reports an error using t.Errorf.
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := colorSubstring(tt.color, tt.substring, tt.text)
            if (err != nil) != tt.shouldErr {
                t.Errorf("colorSubstring() error = %v, shouldErr %v", err, tt.shouldErr)
            }
            if !tt.shouldErr && result != tt.expected {
                t.Errorf("colorSubstring() = %q, want %q", result, tt.expected)
            }
        })
    }
}
