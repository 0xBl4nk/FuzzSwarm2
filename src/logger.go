package src

import (
  "log"
)

// Fatal error and exits the program.
func LogFatal(format string, v ...interface{}) {
  log.Fatalf("[FATAL]"+format, v...)
}

// Informational messages.
func LogInfo(format string, v ...interface{}) {
    log.Printf("[INFO] "+format, v...)
}

// Error messages.
func LogError(format string, v ...interface{}) {
    log.Printf("[ERROR] "+format, v...)
}
