package src

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
    "net/http"
    "net/url"
    "unicode/utf8"
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

// SafeURLEncode safely encodes a value for use in URLs, preventing injection attacks
func SafeURLEncode(value string) string {
    // Validate UTF-8 encoding
    if !utf8.ValidString(value) {
        LogError("Invalid UTF-8 encoding in payload: %s", value)
        return url.QueryEscape(strings.ToValidUTF8(value, ""))
    }
    return url.QueryEscape(value)
}

// SanitizeLogOutput removes potentially sensitive information from log output
func SanitizeLogOutput(input string) string {
    // Remove common authentication patterns from logs
    patterns := []string{
        `password["\s]*[:=]["\s]*[^"\s&]+`,
        `token["\s]*[:=]["\s]*[^"\s&]+`,
        `key["\s]*[:=]["\s]*[^"\s&]+`,
        `secret["\s]*[:=]["\s]*[^"\s&]+`,
    }
    
    result := input
    for _, pattern := range patterns {
        // Replace sensitive data with placeholder
        result = strings.ReplaceAll(result, pattern, "[REDACTED]")
    }
    
    // Limit output length to prevent log flooding
    if len(result) > 500 {
        result = result[:500] + "...[TRUNCATED]"
    }
    
    return result
}

func ApplyHeaders(cfg Config, req *http.Request, fuzzValue string) {
    // Safely encode fuzz value for headers
    safeFuzzValue := SanitizeHeaderValue(fuzzValue)
    
    // Replace 'FUZZ' in headers provided via command line
    if cfg.Headers != "" {
        headers := strings.Split(cfg.Headers, ",")
        for _, header := range headers {
            parts := strings.SplitN(header, ":", 2)
            if len(parts) == 2 {
                key := strings.TrimSpace(parts[0])
                value := strings.TrimSpace(parts[1])
                
                // Validate header key
                if !isValidHeaderName(key) {
                    LogError("Invalid header name: %s", key)
                    continue
                }
                
                // Replace 'FUZZ' placeholder in key and value
                key = strings.ReplaceAll(key, "FUZZ", safeFuzzValue)
                value = strings.ReplaceAll(value, "FUZZ", safeFuzzValue)
                req.Header.Set(key, value)
            } else {
                LogError("Invalid header format: %s", header)
            }
        }
    }

    // Replace 'FUZZ' in headers from headers file
    if len(cfg.HeadersFile) > 0 {
        for key, value := range cfg.HeadersFile {
            // Validate header key
            if !isValidHeaderName(key) {
                LogError("Invalid header name from file: %s", key)
                continue
            }
            
            // Replace 'FUZZ' placeholder in key and value
            key = strings.ReplaceAll(key, "FUZZ", safeFuzzValue)
            value = strings.ReplaceAll(value, "FUZZ", safeFuzzValue)
            req.Header.Set(key, value)
        }
    }

    // Set default Content-Type header if not already set
    if req.Header.Get("Content-Type") == "" {
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    }
}

// SanitizeHeaderValue removes dangerous characters from header values
func SanitizeHeaderValue(value string) string {
    // Remove control characters and normalize
    cleaned := strings.Map(func(r rune) rune {
        if r < 32 || r == 127 {
            return -1 // Remove control characters
        }
        return r
    }, value)
    
    // Limit length to prevent header injection
    if len(cleaned) > 1024 {
        cleaned = cleaned[:1024]
    }
    
    return cleaned
}

// isValidHeaderName checks if a header name is valid according to HTTP specs
func isValidHeaderName(name string) bool {
    if name == "" {
        return false
    }
    
    for _, char := range name {
        if char < 33 || char > 126 || char == ':' {
            return false
        }
    }
    return true
}
