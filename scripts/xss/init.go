package xss

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/0xBl4nk/FuzzSwarm2/src"
)

// LoadXSSPayloads loads XSS payloads and adds them to cfg.Values
func LoadXSSPayloads(cfg *src.Config) {
	const xssPayloadsPath = "scripts/xss/payloads.json"

	payloadTemplates, err := loadXSSPayloads(xssPayloadsPath)
	if err != nil {
		src.LogError("Error loading XSS payloads: %v", err)
		return
	}

	cfg.Values = append(cfg.Values, payloadTemplates...)
	src.LogInfo("XSS payloads added. Total payloads: %d", len(payloadTemplates))
}

// loadXSSPayloads reads the JSON file containing XSS payload templates.
func loadXSSPayloads(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading XSS payloads file: %v", err)
	}

	var payloadTemplates []string
	if err := json.Unmarshal(data, &payloadTemplates); err != nil {
		return nil, fmt.Errorf("error parsing XSS payloads JSON: %v", err)
	}

	return removeDuplicates(payloadTemplates), nil
}

// removeDuplicates removes duplicate entries from a slice of strings.
func removeDuplicates(payloads []string) []string {
	uniqueMap := make(map[string]struct{})
	var uniquePayloads []string

	for _, payload := range payloads {
		if _, exists := uniqueMap[payload]; !exists {
			uniqueMap[payload] = struct{}{}
			uniquePayloads = append(uniquePayloads, payload)
		}
	}

	return uniquePayloads
}