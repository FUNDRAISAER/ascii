package main

import (
    "testing"
)

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
