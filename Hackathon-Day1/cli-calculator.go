package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*CODECRAFTERS — HACKATHON CHALLENGE
     THE CLI CALCULATOR


  Two phases. One calculator. Build it in stages.

  Rules:
  → Write everything in Go. Standard library only.
  → Must compile and run without errors.
  → Push to your the-codecrafters repo in a folder
    named: thecodecrafterthon-day1/
  → Include a README.md explaining how to run it.
  → Complete Phase 1 before opening Phase 2.
  → Deadline: 2 hours from — GO.

   PHASE 1 — GET IT WORKING

  Build a calculator that runs in the terminal,
  accepts user input, and keeps going until
  you decide to stop.

  WHAT TO BUILD

- OPERATIONS
  The user types an expression. The calculator
  responds with the result. Support:

     add <a> <b>   → addition
     sub <a> <b>   → subtraction
     mul <a> <b>   → multiplication
     div <a> <b>   → division

  Example interaction:
     > add 10 4
       ✦ Result: 14

     > mul 3 7
       ✦ Result: 21



-  KEEP IT RUNNING
  The calculator must not stop after one result.
  It waits for the next input, and the next,
  until the user types:

     quit   → exits with a goodbye message

- HELP
  Typing help prints a clean list of all
  supported commands and how to use them.

  VALIDATION — HANDLE ALL OF THESE CLEANLY

  → Division by zero
    must print a clear error. Not a crash.

  → Letters where numbers are expected
    e.g. "add cat dog"
    must print a clear error. Not a crash.

  → Wrong number of arguments
    e.g. "add 5" or "mul 1 2 3"
    must be caught and explained to the user.

  → Unknown commands
    must suggest typing "help".

  → Negative numbers must work perfectly.
    e.g. "add -10 -5  →  -15"

  The golden rule:
  ✦ Nothing the user types should ever
    crash your program. Ever.

  THINK ABOUT

  → How do you read and split user input cleanly?
  → How do you validate before you calculate?
  → What does a helpful error message look like?
  → How do you structure your operations so the
    code doesn't become one giant if-else chain?


PHASE 1 IS DONE WHEN:

  → All 4 operations work correctly.
  → Every bad input is handled — no crashes.
  → quit exits cleanly.
  → help displays usage.
  → Code is committed and pushed.
*/

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- The Calculator ---")

	for {

		fmt.Println("\nPress c to Continue, s for Support, and q to Quit: ")
		output, _ := reader.ReadString('\n')
		output = strings.TrimSpace(output)

		if strings.ToLower(output) == "s" {
			fmt.Println("\nThe Program is built to perform Addition, Subtraction, Multiplication, and Division operations!")
			continue
		} else if strings.ToLower(output) == "q" {
			fmt.Println("\nGoodbye!")
			break
		} else {

			var userinput1, userinput2, operator string

			fmt.Println("\nEnter input1: ")
			fmt.Scan(&userinput1)
			input1, err := strconv.ParseFloat(userinput1, 64)
			if err != nil {
				fmt.Println("\nPermits inputation of numbers only.")
				continue
			}

			fmt.Println("\nEnter operator: \n+, -, *, or /: ")
			fmt.Scan(&operator)

			fmt.Println("\nEnter input2: ")
			fmt.Scan(&userinput2)
			input2, err := strconv.ParseFloat(userinput2, 64)
			if err != nil {
				fmt.Println("\nPermits inputation of numbers only.")
				continue
			}

			switch operator {

			case "+":
				fmt.Println("\nOutput: ", input1+input2)

			case "-":
				fmt.Println("\nOutput: ", input1-input2)

			case "*":
				fmt.Println("\nOutput: ", input1*input2)

			case "/":
				if input1 == 0 || input2 == 0 {
					fmt.Println("\nThe numerator or the denomenator can not be zero!")
				} else {
					fmt.Println("\nOutput: ", input1/input2)
				}
			case "Quit":
				fmt.Println("Goodbye")

			default:
				fmt.Println("\nPermits inputation of only the provided mathematical operators!")
				continue
			}
			if operator == "quit" {
				os.Exit(1)
			}
		}
	}
}
