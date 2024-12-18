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

func safeOrUnsafe(levels [][]int) (int, int) {
	for _, level := range levels {

	}
	return 0, 1
}

func main() {
	var safe int
	var unsafe int

	// Read in file
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("Error occured on file read:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	levels := splitScannerRows(scanner)
	safe, unsafe = safeOrUnsafe(levels)
	println(safe, unsafe)
}
