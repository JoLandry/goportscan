// Package scanner provides functions to perform concurrent TCP port scanning on a given IP address.
// It can check reachability of the host, scan a specified port range, and output results in JSON format.
package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type ScanResult struct {
	Port int  `json:"port"`
	Open bool `json:"open"`
}

// ScanPortsInRange scans ports in the range [start, end] on the given IP address.
// It sends ScanResult structs based on the scan into the results channel.
// It is intended to be run as a goroutine.
// The given WaitGroup is used to track when scanning is complete.
func ScanPortsInRange(start int, end int, ip string, wg *sync.WaitGroup, results chan<- ScanResult) {
	defer wg.Done()
	for currPort := start; currPort <= end; currPort++ {
		address := fmt.Sprintf("%s:%d", ip, currPort)
		conn, err := net.DialTimeout("tcp", address, 300*time.Millisecond)
		if err == nil {
			results <- ScanResult{currPort, true}
			conn.Close()
		} else {
			results <- ScanResult{currPort, false}
		}
	}
}

// ScanPortsAndFormatJSON performs concurrent scanning of ports between start and end on a given IP.
// It uses 250 worker goroutines and writes the scan results as JSON to the specified output path.
// Returns an error if writing to JSON fails.
func ScanPortsAndFormatJSON(start int, end int, ip string, outputPath string) error {
	const nbWorkers = 250
	var wg sync.WaitGroup

	// Declaring the buffered channels
	results := make(chan ScanResult, 1000)

	nbPortsPerRoutine := (end - start + 1) / nbWorkers
	for i := 0; i < nbWorkers; i++ {
		// Compute ports range
		from := start + i*nbPortsPerRoutine
		to := from + nbPortsPerRoutine - 1

		if i == nbWorkers-1 {
			to = end
		}
		wg.Add(1)
		go ScanPortsInRange(from, to, ip, &wg, results)
	}

	// Collect results
	allResults := collectResults(results, &wg)

	// Write results to file
	if err := WriteResultsToJSON(allResults, outputPath); err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}

// collectResults reads all ScanResult values from the results channel after all workers are done.
// It returns a slice of ScanResult structs.
func collectResults(results chan ScanResult, wg *sync.WaitGroup) []ScanResult {
	// Close channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	var all []ScanResult
	for res := range results {
		all = append(all, res)
	}
	return all
}
