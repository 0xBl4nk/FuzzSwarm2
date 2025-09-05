package src

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	"context"
	"sync/atomic"
	"net/url"
)


func StartFuzzing(cfg Config) {
  LogInfo("Start fuzzing...")
  LogInfo("Using %d threads", cfg.Threads)
  LogInfo("Total values to fuzz: %d", len(cfg.Values))

  // Create context for graceful cancellation
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  var wg sync.WaitGroup
  semaphore := make(chan struct{}, cfg.Threads)
  client := CreateClient(cfg.UseProxy, cfg.Timeout, cfg.UseSSL, cfg.FollowRedirects)

  // Progress tracking
  var completed int64
  total := int64(len(cfg.Values))
  
  // Progress reporter goroutine
  go func() {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()
    
    for {
      select {
      case <-ticker.C:
        current := atomic.LoadInt64(&completed)
        if current > 0 {
          percentage := float64(current) / float64(total) * 100
          LogInfo("Progress: %d/%d (%.1f%%) completed", current, total, percentage)
        }
      case <-ctx.Done():
        return
      }
    }
  }()

  for _, value := range cfg.Values {
    select {
    case <-ctx.Done():
      LogInfo("Fuzzing cancelled")
      return
    case semaphore <- struct{}{}:
      wg.Add(1)
      go func (val string) {
        defer func() {
          atomic.AddInt64(&completed, 1)
          wg.Done()
          <-semaphore
        }()
        FuzzRequest(cfg, client, val, ctx)
      }(value)
    }
  }
  
  wg.Wait()
  LogInfo("Fuzzing completed. Total requests: %d", total)
}

func FuzzRequest(cfg Config, client *http.Client, value string, ctx context.Context) {
  var placehold = "FUZZ"
  
  // Safely encode the value for URL usage
  encodedValue := SafeURLEncode(value)
  requestURL := strings.Replace(cfg.URL, placehold, encodedValue, -1)
  
  // Validate the final URL
  if _, err := url.Parse(requestURL); err != nil {
    LogError("Invalid URL after fuzzing for value '%s': %v", value, err)
    return
  }
  
  var req *http.Request
  var err error

  if cfg.Method == "POST" || cfg.Method == "PUT" || cfg.Method == "PATCH" {
    fuzzedData := strings.ReplaceAll(cfg.Data, placehold, value)
    req, err = http.NewRequestWithContext(ctx, cfg.Method, requestURL, strings.NewReader(fuzzedData))
    if err != nil {
      LogError("Failed to create %s request for value '%s': %v", cfg.Method, value, err)
      return
    }

  } else {
    req, err = http.NewRequestWithContext(ctx, cfg.Method, requestURL, nil)
    if err != nil {
      LogError("Failed to create %s request for value '%s': %v", cfg.Method, value, err)
      return
    }
  }

  // Set random User-Agent for stealth
  req.Header.Set("User-Agent", GetRandomUserAgent())
  
  // Apply custom headers
  ApplyHeaders(cfg, req, value) 

  var resp *http.Response
  for attempt := 1; attempt <= cfg.Retries; attempt++ {
    select {
    case <-ctx.Done():
      return
    default:
    }
    
    if cfg.RateLimit > 0 {
      time.Sleep(time.Millisecond * time.Duration(cfg.RateLimit))
    }

    resp, err = client.Do(req)
    if err == nil {
      break
    }
    LogError("Request failed for value '%s' on attempt %d: %v", value, attempt, err)
    
    // Exponential backoff
    backoffTime := time.Duration(attempt*attempt) * time.Second
    if backoffTime > 30*time.Second {
      backoffTime = 30 * time.Second
    }
    
    select {
    case <-time.After(backoffTime):
    case <-ctx.Done():
      return
    }
  }

  if err != nil {
    LogError("All retry attempts failed for value '%s': %v", value, err)
    return
  }
  defer resp.Body.Close()

  bodyBytes, err := io.ReadAll(resp.Body)
  if err != nil {
    LogError("Failed to read response body for value '%s': %v", value, err)
    return
  }

  responseBody := string(bodyBytes)
  responseSize := len(responseBody)

  if cfg.FilterSize > 0 && responseSize == cfg.FilterSize {
    return
  }

  printResponse(cfg, value, resp.StatusCode, responseSize, responseBody)
}

func printResponse(cfg Config, value string, statusCode int, responseSize int, responseBody string) {
  // Sanitize output for security
  sanitizedValue := SanitizeLogOutput(value)
  sanitizedBody := SanitizeLogOutput(responseBody)

  result := Result{
    Value:        sanitizedValue,
    StatusCode:   statusCode,
    ResponseSize: responseSize,
    URL:          cfg.URL,
    Method:       cfg.Method,
    Timestamp:    time.Now().Format("2006-01-02T15:04:05Z"),
  }
  
  if cfg.Verbose {
    previewLength := 100
    if len(sanitizedBody) > previewLength {
      result.Preview = sanitizedBody[:previewLength] + "..."
    } else {
      result.Preview = sanitizedBody
    }
  }
  
  LogResult(cfg, result)
}


