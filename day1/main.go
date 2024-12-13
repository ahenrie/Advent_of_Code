package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	first_column := []int{}
	second_column := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if num, err := strconv.Atoi(parts[0]); err == nil {
			first_column = append(first_column, num)
		}
		if num, err := strconv.Atoi(parts[1]); err == nil {
			second_column = append(second_column, num)
		}
	}

	sort.Ints(first_column)
	sort.Ints(second_column)
	fmt.Println(first_column)
	fmt.Println(second_column)

	var total int
	for i, _ := range first_column {
		diff := first_column[i] - second_column[i]
		total += int(math.Abs(float64(diff)))

	}
	fmt.Println(total)
}
