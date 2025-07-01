// Package scanner provides functions to perform concurrent TCP port scanning on a given IP address.
// It can check reachability of the host, scan a specified port range, and output results in JSON format.
package scanner

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteResultsToJSON marshals the given slice of ScanResult into formatted JSON
// and writes it to the specified file path.
// Returns an error if marshalling or file writing fails.
func WriteResultsToJSON(results []ScanResult, outputPath string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
