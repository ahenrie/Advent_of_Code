package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printLine(inScanner *bufio.Scanner) {
	for inScanner.Scan() {
		fmt.Println(inScanner.Text())
	}
}

func splitScannerRows(inScanner *bufio.Scanner) [][]int {
	slices := [][]int{}
	for inScanner.Scan() {
		buffer := []int{}
		line := inScanner.Text()
		// Split the line into parts by whitespace
		parts := strings.Fields(line)
		for _, part := range parts {
			if num, err := strconv.Atoi(part); err == nil {
				//fmt.Println(num)
				buffer = append(buffer, num)
			} else {
				fmt.Println("Error converting string to int:", err)
			}
		}
		slices = append(slices, buffer)
	}
	return slices
}

// increasing or decreasing
func increasingOrDecreasing(slice []int) string {
	if len(slice) < 2 {
		return "neither"
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[i-1] {
			decreasing = false
		}
		if slice[i] < slice[i-1] {
			increasing = false
		}
	}

	if increasing {
		return "increasing"
	}
	if decreasing {
		return "decreasing"
	}
	return "neither"
}

// adjacent checks if the difference between adjacent numbers in the slice is 1, 2, or 3.
func adjacent(slice []int) bool {
	// I do not want to convert to float64
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 1; i < len(slice); i++ {
		diff := abs(slice[i] - slice[i-1])
		if diff != 1 && diff != 2 && diff != 3 {
			return false
		}
	}
	return true
}

func safeOrUnsafe(levels [][]int) (int, int) {
	safe := 0
	unsafe := 0
	for _, level := range levels {
		if increasingOrDecreasing(level) != "neither" && adjacent(level) == true {
			safe = safe + 1
		} else {
			unsafe = unsafe + 1
		}
	}
	return safe, unsafe
}

func main() {
	// Read in file
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("Error occured on file read:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	levels := splitScannerRows(scanner)
	safe, unsafe := safeOrUnsafe(levels)
	fmt.Printf("Safe: %v, Unsafe: %v:", safe, unsafe)
}
