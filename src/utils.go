package src

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// ReadHeaders reads an HTTP headers file, ignoring empty lines and comments.
func ReadHeaders(path string) (map[string]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    headers := make(map[string]string)
    lineNumber := 0
    for scanner.Scan() {
        lineNumber++
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue // Ignore empty lines and comments
        }
        parts := strings.SplitN(line, ": ", 2)
        if len(parts) == 2 {
            headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
        } else {
            return nil, fmt.Errorf("invalid header format at line %d: %s", lineNumber, line)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return headers, nil
}

// ReadValues reads a wordlist file for fuzzing, ignoring empty lines and comments.
func ReadValues(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var values []string
    scanner := bufio.NewScanner(file)
    lineNumber := 0
    for scanner.Scan() {
        lineNumber++
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        values = append(values, line)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return values, nil
}
