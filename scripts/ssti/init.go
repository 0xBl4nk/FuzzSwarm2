package ssti

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/0xBl4nk/FuzzSwarm2/src"
)

// LoadSSTIPayloads loads SSTI payloads and adds them to cfg.Values
func LoadSSTIPayloads(cfg *src.Config) {
	const sstiPayloadsPath = "scripts/ssti/payloads.json" // Fixed path to payloads.json

	payloadTemplates, err := loadSSTIPayloads(sstiPayloadsPath)
	if err != nil {
		src.LogError("Error loading SSTI payloads: %v", err)
		return
	}

	// Validate payloads
	if err := validatePayloads(payloadTemplates); err != nil {
		src.LogError("Error validating SSTI payloads: %v", err)
		return
	}

	uniquePayloads := removeDuplicates(payloadTemplates)
	sstiValues := generateSSTIValues(uniquePayloads) // Generates SSTI payloads without duplication
	cfg.Values = append(cfg.Values, sstiValues...)
	src.LogInfo("SSTI payloads added. Total payloads: %d", len(sstiValues))
}

// loadSSTIPayloads reads the JSON file containing SSTI payload templates.
func loadSSTIPayloads(path string) ([]string, error) {
	data, err := os.ReadFile(path) // Replaces ioutil.ReadFile with os.ReadFile
	if err != nil {
		return nil, fmt.Errorf("error reading SSTI payloads file: %v", err)
	}

	var payloadTemplates []string
	if err := json.Unmarshal(data, &payloadTemplates); err != nil {
		return nil, fmt.Errorf("error parsing SSTI payloads JSON: %v", err)
	}

	return payloadTemplates, nil
}

// validatePayloads checks if all payloads contain the 'UNIQUE_NUM' placeholder.
func validatePayloads(payloads []string) error {
	for _, payload := range payloads {
		if !strings.Contains(payload, "UNIQUE_NUM") {
			return fmt.Errorf("payload '%s' does not contain the 'UNIQUE_NUM' placeholder", payload)
		}
	}
	return nil
}

// generateSSTIValues generates payloads by replacing UNIQUE_NUM with a unique number for each payload.
func generateSSTIValues(payloadTemplates []string) []string {
	var values []string

	uniqueNum := generateUniqueNum() // Generates a unique number for all payloads
	for _, template := range payloadTemplates {
		payload := strings.ReplaceAll(template, "UNIQUE_NUM", fmt.Sprintf("%d", uniqueNum))
		values = append(values, payload)
	}

	return values
}

// generateUniqueNum generates a unique number based on the current time.
func generateUniqueNum() int {
	// Uses the current timestamp in milliseconds to ensure uniqueness
	return int(time.Now().UnixNano() / int64(time.Millisecond))
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
