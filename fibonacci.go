package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/John-Hatton/cmdline_go"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const version = "1.1.0"

var outputBuffer bytes.Buffer

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
		outputBuffer.WriteString("Fibonacci sequence: \n")
		fmt.Print("Fibonacci sequence: \n")
		for i := uint64(2); i <= n; i++ {
			if memo[i] != big.NewInt(1) && memo[i] != big.NewInt(0) {
				fibStr := fmt.Sprintf("%d\n", memo[i])
				outputBuffer.WriteString(fibStr)
				fmt.Printf("%d\n", memo[i])
			}
		}
		outputBuffer.WriteString("\n")
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
	cmdLine.HelpText = "Usage: fibonacci_go [OPTIONS]\n\nOptions:\n  -d, --debug        Set DEBUG flag true\n  -v, --version      Print version number\n  -h, --help         Print this table\n  -f FILENAME        Take a file in as input\n  -i INPUT_STRING    Process an input string\n  -o LOG_TO_CONSOLE  Log output to console\n  -l LOG_FILENAME    Save output to log file\n"

	myVer := fmt.Sprintf("Fibonacci_Go -- Version: %s\nby: John Hatton", version)
	cmdLine.VersionText = myVer

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
		outputBuffer.WriteString("Debug mode enabled \n")
		fmt.Println("Debug mode enabled")
	}

	// If the user provided a filename with the"-f" flag, read the input from the file and set it as the value of input variable.

	if cmdLine.FileName != "" {
		file, err := os.Open(cmdLine.FileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		input := scanner.Text()

		// Call the Fibonacci function with the parsed input value and debug flag.
		result := Fibonacci(parseInput(input), cmdLine.Debug, make(map[uint64]*big.Int))

		// If the user provided the "-o" flag, print the result to the console.
		if cmdLine.LogToConsole {
			fmt.Println(result)
		}

		// If the user provided the "-l" flag, save the result to the specified log file.
		if cmdLine.LogFileName != "" {
			file, err := os.OpenFile(cmdLine.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()

			_, err = fmt.Fprintf(file, "%d\n", result)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	} else if cmdLine.InputText != "" {
		// If the user provided an input string with the "-i" flag, call the Fibonacci function with that value and debug flag.
		result := Fibonacci(parseInput(cmdLine.InputText), cmdLine.Debug, make(map[uint64]*big.Int))

		// Print no matter what, if we get input
		fibStr := fmt.Sprintf("Fibonacci(%d) = %d\n", parseInput(cmdLine.InputText), result)
		outputBuffer.WriteString(fibStr)
		fmt.Printf("Fibonacci(%d) = %d\n", parseInput(cmdLine.InputText), result)

		// If the user provided the "-o" flag, print the result to the console.
		if cmdLine.Output {

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		//If the user provided the "-l" flag, save the result to the specified log file.
		if cmdLine.LogFileName != "" {
			file, err := os.OpenFile(cmdLine.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()

			_, err = fmt.Fprintf(file, "%s", outputBuffer.String())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	} else {
		// If no input was provided, print the help text.
		fmt.Println(cmdLine.HelpText)
	}
}

// Function to parse the input value to an unsigned 64-bit integer.
func parseInput(input string) uint64 {
	n, err := strconv.ParseUint(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return n
}
