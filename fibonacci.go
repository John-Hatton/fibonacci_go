package main

import (
	"bufio"
	"fmt"
	"github.com/John-Hatton/cmdline_go"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const version = "1.0.6"

func Fibonacci(n uint64, debug bool, memo map[uint64]*big.Int) *big.Int {
	// Check if the value of `n` has been already computed and saved in the map.
	if val, ok := memo[n]; ok {
		return val
	}

	// If `n` is less than 2, return `n`.
	if n < 2 {
		return big.NewInt(int64(n))
	}

	// Calculate the Fibonacci number using a loop and memoize it in the map.
	a, b := big.NewInt(0), big.NewInt(1)
	for i := uint64(2); i <= n; i++ {
		a, b = b, new(big.Int).Add(a, b)
		memo[i] = new(big.Int).Set(b)
	}

	// If `debug` is true, print the sequence with a debug message.
	if debug {
		fmt.Print("Fibonacci sequence: ")
		for i := uint64(0); i <= n; i++ {
			fmt.Printf("%d\n", memo[i])
		}
		fmt.Println()
	}

	return b
}

func main() {
	// Get command-line arguments using os.Args.
	args := os.Args[1:]

	// Create an instance of CommandLine struct from the cmdline_go package.
	cmdLine := cmdline_go.CommandLine{}

	// Set the help text for the command-line interface.
	cmdLine.HelpText = "Usage: fibonacci_go [OPTIONS]\n\nOptions:\n  -d, -debug      Set DEBUG flag true\n  -v, -version    Print version number\n  -h, -help       Print this table\n  -f FILENAME     Print report about file\n  -i INPUT_STRING Process an input string\n"

	mystr := fmt.Sprintf("Fibonacci_Go -- Version: %s\nby: John Hatton", version)
	cmdLine.VersionText = mystr

	// Parse the command-line arguments using the Parse method of CommandLine struct.
	err := cmdLine.Parse(args)
	if err != nil {
		// Print an error message and exit if there is an error in parsing the arguments.
		fmt.Println(err)
		os.Exit(1)
	}

	// Process the parsed command-line arguments using the Process method of CommandLine struct.
	err = cmdLine.Process()
	if err != nil {
		// Print an error message and exit if there is an error in processing the arguments.
		fmt.Println(err)
		os.Exit(1)
	}

	// If the debug flag is set to true, print a message indicating that debug mode is enabled.
	if cmdLine.Debug {
		fmt.Println("Debug mode enabled")
	}

	// If the user provided a filename with the -f option, read the input file and call the Fibonacci function with the values read from the file.
	if cmdLine.FileName != "" {
		// Open the input file.
		file, err := os.Open(cmdLine.FileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		var n uint64
		// Use a scanner to read the file line by line.
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			// Skip comment lines that start with #.
			if strings.HasPrefix(text, "#") {
				continue
			}
			// Convert the text to an integer and call the Fibonacci function with it.
			parsedN, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				fmt.Printf("Error converting %s to integer: %v\n", text, err)
				continue
			}
			n = parsedN
			result := Fibonacci(n, cmdLine.Debug, make(map[uint64]*big.Int))
			fmt.Printf("Fibonacci(%d) = %d\n", n, result)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	} else if cmdLine.InputText != "" {
		// Read input text
		// Call Fibonacci with the value read from the input text
		lines := strings.Split(cmdLine.InputText, "\n")
		var n uint64
		for _, line := range lines {
			if strings.HasPrefix(line, "#") {
				continue
			}
			// Convert the text to a uint64 and call the Fibonacci function with it.
			n, err = strconv.ParseUint(line, 10, 64)
			if err != nil {
				fmt.Printf("Error converting %s to uint64: %v\n", line, err)
				continue
			}
			result := Fibonacci(n, cmdLine.Debug, make(map[uint64]*big.Int))
			fmt.Printf("Fibonacci(%d) = %d\n", n, result)
		}
	} else {
		// Read user input from the command line
		var n uint64
		fmt.Print("Enter a number: ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println("Invalid input")
			os.Exit(1)
		}
		result := Fibonacci(n, cmdLine.Debug, make(map[uint64]*big.Int))
		fmt.Printf("Fibonacci(%d) = %d\n", n, result)
	}

	if cmdLine.Help {
		cmdLine.PrintHelp()
		os.Exit(0)
	}
	if cmdLine.Version {
		cmdLine.PrintVersion()
		os.Exit(0)
	}
}
