package main

import (
	"fmt"
	"os"

	"github.com/JoLandry/goportscan/config"
	"github.com/JoLandry/goportscan/scanner"
)

func main() {
	fmt.Println("Hello World !")
	fmt.Println("Welcome to goportscan !")
	fmt.Println()

	// Parse CLI options
	options, err := config.ParseOptions()
	fmt.Println()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		config.PrintHelp()
		os.Exit(1)
	}

	// Options
	start := options.StartPort
	end := options.EndPort
	outputPath := options.OutputPath
	hostIP := options.IP

	// Launch the algorithm
	scanner.ScanPortsAndFormatJSON(start, end, hostIP, outputPath)

	fmt.Println("Finished scanning, go take a look at file : " + outputPath)
}
