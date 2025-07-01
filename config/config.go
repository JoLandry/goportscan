// Package config provides features to parse and validate CLI options
package config

import (
	"fmt"
	"os"
	"strconv"
)

type Options struct {
	IP         string
	StartPort  int
	EndPort    int
	OutputPath string
}

// ParseOptions parses command-line arguments from os.Args,
// extracts options for IP address, start and end port, and output file path.
//
// Expected args format:
//
//	[]string{
//	  "goportscan",
//	  "-ip", "192.168.0.1",
//	  "-start", "20",
//	  "-end", "1000",
//	  "-o", "output.json",
//	}
//
// Recognized options:
//
//	-ip       IP address of the host to scan (required)
//	-start    Starting port number (required)
//	-end      Ending port number (required)
//	-o        Output file path for JSON results (required)
//	-h, --help  Prints this help message and exits
//
// Returns an Options struct with parsed values or an error if parsing fails.
func ParseOptions() (Options, error) {
	// Retrieve args but skip binary
	args := os.Args[1:]
	options := Options{}
	var err error

	for i := 0; i < len(args); i++ {
		switch args[i] {
		// IP
		case "-ip":
			i++
			if i < len(args) {
				options.IP = args[i]
			} else {
				return options, fmt.Errorf("missing value for -ip")
			}
		// Starting port
		case "-start":
			i++
			if i < len(args) {
				options.StartPort, err = strconv.Atoi(args[i])
				if err != nil {
					return options, fmt.Errorf("invalid value for -start: %v", err)
				}
			} else {
				return options, fmt.Errorf("missing value for -start")
			}
		// Ending port
		case "-end":
			i++
			if i < len(args) {
				options.EndPort, err = strconv.Atoi(args[i])
				if err != nil {
					return options, fmt.Errorf("invalid value for -end: %v", err)
				}
			} else {
				return options, fmt.Errorf("missing value for -end")
			}
		// Output file path
		case "-o":
			i++
			if i < len(args) {
				options.OutputPath = args[i]
			} else {
				return options, fmt.Errorf("missing value for -o")
			}
		// Ask for help
		case "-h", "--help":
			PrintHelp()
			os.Exit(0)
		// Unknown option
		default:
			return options, fmt.Errorf("unknown option: %s", args[i])
		}
	}

	// Check that options were given
	if options.IP == "" || options.StartPort == 0 || options.EndPort == 0 || options.OutputPath == "" {
		return options, fmt.Errorf("missing one or more required options")
	}

	return options, nil
}

// PrintHelp prints usage instructions and options description to standard output.
func PrintHelp() {
	fmt.Print(`Usage: goportscan -ip <IP> -start <startPort> -end <endPort> -o <output.json>
Options:
  -ip        IP address of the host to scan
  -start     Starting port number
  -end       Ending port number
  -o         Output file path (JSON format)
  -h, --help Show this help message
`)
}
