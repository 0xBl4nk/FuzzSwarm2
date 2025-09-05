package sqli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/0xBl4nk/FuzzSwarm2/src"
)

// LoadSQLiPayloads loads SQL injection payloads and adds them to cfg.Values
func LoadSQLiPayloads(cfg *src.Config) {
	const sqliPayloadsPath = "scripts/sqli/payloads.json"

	payloadTemplates, err := loadSQLiPayloads(sqliPayloadsPath)
	if err != nil {
		src.LogError("Error loading SQLi payloads: %v", err)
		return
	}

	// No need for unique numbers in SQL injection, just use payloads directly
	cfg.Values = append(cfg.Values, payloadTemplates...)
	src.LogInfo("SQL injection payloads added. Total payloads: %d", len(payloadTemplates))
}

// loadSQLiPayloads reads the JSON file containing SQL injection payload templates.
func loadSQLiPayloads(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading SQLi payloads file: %v", err)
	}

	var payloadTemplates []string
	if err := json.Unmarshal(data, &payloadTemplates); err != nil {
		return nil, fmt.Errorf("error parsing SQLi payloads JSON: %v", err)
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