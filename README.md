# goportscan

`goportscan` is a small concurrent TCP port scanner written in Go. It scans a specified range of ports on a target IP address and outputs the results in JSON format. This CLI tool is designed for ease of use, fast scanning with concurrency, and output options.

## Version

1.0.0

---

## Features

- Scan TCP ports on a given IP address
- Specify start and end ports for scanning range
- Concurrent scanning with configurable worker goroutines
- Output scan results as JSON file
- Command-line interface with help and validation

---

## Usage

Example of usage :

```bash
go run main.go -ip <IP> -start <startPort> -end <endPort> -o <output.json>
```

You should replace these arguments with the actual values you want to run the program with.

---

## Example output

```json
[
  {
    "port": 25,
    "open": false
  },
  {
    "port": 26,
    "open": false
  },
  {
    "port": 27,
    "open": false
  },
  {
    "port": 28,
    "open": false
  },
  {
    "port": 29,
    "open": false
  },
  {
    "port": 30,
    "open": false
  }
]
```

---

## Installation

```bash
go get github.com/JoLandry/goportscan
```

Then build or run your scanner (from the root of the project):

```bash
go build ./...
./goportscan -ip <IP> -start <startPort> -end <endPort> -o <output.json>
```

Or you could also run (from the root of the project):

```bash
go build ./...
go run main.go -ip <IP> -start <startPort> -end <endPort> -o <output.json>
```

---

## License

This project was created as a personal initiative, outside any official academic coursework, during my Master's in Computer Science at the University of Bordeaux.

It is released under the MIT License.

You are free to:
    use, copy, modify, and distribute this code,
    as long as you retain the copyright notice.

This project is provided "as is", without any warranty of any kind.

© 2025 Landry Jonathan