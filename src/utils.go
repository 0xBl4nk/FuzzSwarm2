package src

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
    "net/http"
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
func ReadWordlistValues(path string) ([]string, error) {
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

// Generates a slice of strings based on the provided range string
func parseRange(rangeStr string) ([]string, error) {
    parts := strings.Split(rangeStr, ",")
    if len(parts) != 2 {
        return nil, fmt.Errorf("range format should be start-end,digits (e.g., 1-10000,3)")
    }
    rangeParts := strings.Split(parts[0], "-")
    if len(rangeParts) != 2 {
        return nil, fmt.Errorf("range bounds format should be start-end")
    }
    start, err := strconv.Atoi(rangeParts[0])
    if err != nil {
        return nil, fmt.Errorf("invalid start value: %v", err)
    }
    end, err := strconv.Atoi(rangeParts[1])
    if err != nil {
        return nil, fmt.Errorf("invalid end value: %v", err)
    }
    digits, err := strconv.Atoi(parts[1])
    if err != nil {
        return nil, fmt.Errorf("invalid digits value: %v", err)
    }

    var values []string
    for i := start; i <= end; i++ {
        values = append(values, fmt.Sprintf("%0*d", digits, i))
    }
    return values, nil
  }

  func ApplyHeaders(cfg Config, req *http.Request) {
    if cfg.Headers != "" {
        headers := strings.Split(cfg.Headers, ",")
        
        for _, header := range headers {
            parts := strings.SplitN(header, ":", 2)
            if len(parts) == 2 {
                key := strings.TrimSpace(parts[0])
                value := strings.TrimSpace(parts[1])
                req.Header.Set(key, value)
            } else {
                LogError("Invalid header format: %s", header)
            }
        }
    }

    if req.Header.Get("Content-Type") == "" {
        req.Header.Set("Content-Type", "application/json")
    }
}
