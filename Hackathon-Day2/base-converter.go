package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*THE BASE CONVERTER
 Concept: Number Systems & strconv

Rules:
 → DO NOT USE GOOGLE OR ANY AI TOOL; DO IT YOURSELF - you can share ideas with others, but not copy code.
 → Write everything in Go. Standard library only.
 → Must compile and run without errors.
 → Push to your the-codecrafters repo in a folder
   named: thecodecrafterthon-day2/
 → Include a README.md explaining how to run it

 go-reloaded needs to convert hex and binary
 strings into decimal numbers. This project
 teaches you exactly that — and makes you think
 about what happens when the input is garbage.

 Build a CLI tool that converts numbers between
 bases. It runs from the terminal like this:

    > convert 1E hex
      ✦ Decimal: 30

    > convert 10 bin
      ✦ Decimal: 2

    > convert 255 dec
      ✦ Binary:  11111111
      ✦ Hex:     FF

 Requirements:

 → Support three input bases: hex, bin, dec.
 → For dec input, output both binary and hex.
 → For hex and bin input, output only decimal.
 → All hex output must be uppercase.
 → The program keeps running until: quit

 Validation — handle all of these cleanly:
 → "1G" is not valid hex.
 → "10201" is not valid binary.
 → "abc" is not a valid decimal.
 → Negative numbers must be supported for dec.
 → Empty input must not crash the program.

 Think about:
 → Which strconv functions handle base
   conversions for you?
 → How do you detect which characters are
   valid for a given base?
 → What is the difference between ParseInt
   and ParseUint — and does it matter here?

 ✦ Nail this and (hex) / (bin) in go-reloaded
   becomes a 5-minute task.
*/

func hexToDecimal(text string) (string, error) {
	hex, err := strconv.ParseInt(text, 16, 64)
	if err != nil {
		return "\nInvalid hex", err
	}
	return fmt.Sprintln(hex), nil
}

func binToDecimal(str string) (string, error) {
	bin, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		return "\nInvalid bin", err
	}
	return fmt.Sprintln(bin), nil
}

func decTobinhex(text string) (string, string, error) {
	allInt, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return "\nInvalid hex", "", err
	}
	final := strconv.FormatInt(allInt, 16)

	return fmt.Sprintf("%v", strings.ToUpper(final)), fmt.Sprintf("%08b", allInt), nil

}

func main() {
start:
	userinput := bufio.NewScanner(os.Stdin)
	fmt.Println("\nEnter any of the following commands to perform a Base Convertion: 'convert 1E hex', 'convert 10 bin', 'convert 255 dec'; and/or press 'q' to quit: ")
	if !userinput.Scan() {
		fmt.Println("Invalid")
		goto start
	}

	input := userinput.Text()
	fields := strings.Fields(input)

	if input == "q" {
		fmt.Println("\nGoodbye!")
		return
	}

	if input == "" || len(fields) < 3 || len(fields) > 3 {
		fmt.Println("\nInvalid arguments. Try again!")
		goto start
	}

	switch fields[2] {
	case "hex":
		result, err := hexToDecimal(fields[1])
		if err != nil {
			fmt.Printf("'%v' is not valid hex\n", fields[1])
			goto start
		}
		fmt.Printf("\nDecimal: %v\n", result)
		goto start

	case "bin":
		result, err := binToDecimal(fields[1])
		if err != nil {
			fmt.Printf("'%v' is not valid binary\n", fields[1])
			goto start
		}
		fmt.Printf("\nBinary: %v\n", result)
		goto start

	case "dec":
		result1, result2, err := decTobinhex(fields[1])
		if err != nil {
			fmt.Printf("'%v' is not a valid decimal\n", fields[1])
			goto start
		}
		fmt.Printf("\nBinary: %v\n", result2)
		fmt.Printf("Hex: %v\n", result1)
		goto start

	default:
		fmt.Println("\nInvalid arguments. Try again!")
		goto start
	}
}
