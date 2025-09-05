package src

import (
  "log"
  "fmt"
  "github.com/fatih/color"
)

// LogLevel represents different logging levels
type LogLevel int

const (
  DEBUG LogLevel = iota
  INFO
  WARN
  ERROR
  FATAL
)

var currentLogLevel = INFO

// SetLogLevel sets the current logging level
func SetLogLevel(level LogLevel) {
  currentLogLevel = level
}

// Fatal error and exits the program.
func LogFatal(format string, v ...interface{}) {
  log.Fatalf("[FATAL] "+format, v...)
}

// Informational messages.
func LogInfo(format string, v ...interface{}) {
  if currentLogLevel <= INFO {
    log.Printf("[INFO] "+format, v...)
  }
}

// Warning messages.
func LogWarn(format string, v ...interface{}) {
  if currentLogLevel <= WARN {
    log.Printf("[WARN] "+format, v...)
  }
}

// Error messages.
func LogError(format string, v ...interface{}) {
  if currentLogLevel <= ERROR {
    log.Printf("[ERROR] "+format, v...)
  }
}

// Debug messages.
func LogDebug(format string, v ...interface{}) {
  if currentLogLevel <= DEBUG {
    log.Printf("[DEBUG] "+format, v...)
  }
}

// Result represents a fuzzing result for structured output
type Result struct {
  Value        string `json:"value"`
  StatusCode   int    `json:"status_code"`
  ResponseSize int    `json:"response_size"`
  URL          string `json:"url"`
  Method       string `json:"method"`
  Timestamp    string `json:"timestamp"`
  Preview      string `json:"preview,omitempty"`
}

// LogResult logs a result in the specified format
func LogResult(cfg Config, result Result) {
  switch cfg.OutputFormat {
  case "json":
    logResultJSON(result)
  case "csv":
    logResultCSV(result)
  default:
    logResultText(cfg, result)
  }
}

func logResultJSON(result Result) {
  // This would be implemented with proper JSON marshaling
  fmt.Printf(`{"value":"%s","status_code":%d,"response_size":%d,"url":"%s","method":"%s","timestamp":"%s"}` + "\n",
    result.Value, result.StatusCode, result.ResponseSize, result.URL, result.Method, result.Timestamp)
}

func logResultCSV(result Result) {
  fmt.Printf("%s,%d,%d,%s,%s,%s\n",
    result.Value, result.StatusCode, result.ResponseSize, result.URL, result.Method, result.Timestamp)
}

func logResultText(cfg Config, result Result) {
  // Use existing color-coded output
  colorFunc := getColorFunc(result.StatusCode)
  if cfg.Verbose && result.Preview != "" {
    colorFunc.Printf("Value: %s [%d] - Response size: %d - Preview: %s\n", 
      result.Value, result.StatusCode, result.ResponseSize, result.Preview)
  } else {
    colorFunc.Printf("Value: %s [%d] - Response size: %d\n", 
      result.Value, result.StatusCode, result.ResponseSize)
  }
}

func getColorFunc(statusCode int) *color.Color {
  switch {
  case statusCode >= 200 && statusCode < 300:
      return color.New(color.FgGreen)
  case statusCode >= 300 && statusCode < 400:
    return color.New(color.FgYellow)
  case statusCode >= 400 && statusCode < 500:
    return color.New(color.FgRed)
  case statusCode >= 500:
    return color.New(color.FgMagenta)
  default:
    return color.New(color.FgWhite)
  }
}
