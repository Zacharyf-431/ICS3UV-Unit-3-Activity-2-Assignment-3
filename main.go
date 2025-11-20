/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-20
	* This file calculates cents 
	*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // INPUT
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the amount of cents: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    totalCents, _ := strconv.Atoi(input)

    // PROCESS
    centsLeft := totalCents

		// calculates the amount of toonies 
    toonies := (centsLeft - (centsLeft % 200)) / 200
    centsLeft = centsLeft % 200

		// calculates the amount of loonies 
    loonies := (centsLeft - (centsLeft % 100)) / 100
    centsLeft = centsLeft % 100

		// calculates the amount of quarters
    quarters := (centsLeft - (centsLeft % 25)) / 25
    centsLeft = centsLeft % 25

		// calculates the amount of dimes
    dimes := (centsLeft - (centsLeft % 10)) / 10
    centsLeft = centsLeft % 10

		// calculates the amount of nickels
    nickels := (centsLeft - (centsLeft % 5)) / 5
    centsLeft = centsLeft % 5

    // OUTPUT
    fmt.Printf("Your change is: %d toonies, %d dollar, %d quarters, %d dimes, %d nickels and %d cents\n", toonies, loonies, quarters, dimes, nickels, centsLeft)
		fmt.Println("\nDone.")
}