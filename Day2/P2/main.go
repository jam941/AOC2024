package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func isValid(report []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
		if diff < 0 {
			increasing = false
		}
		if diff > 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func isSafePart2Check(report []int) bool {
	for i := 0; i < len(report); i++ {
		modified := append([]int{}, report[:i]...)
		modified = append(modified, report[i+1:]...)
		if isValid(modified) {
			return true
		}
	}
	return false
}

func countSafeReportsWithDampener(reports [][]int) int {
	safeCount := 0

	for _, report := range reports {
		if isValid(report) || isSafePart2Check(report) {
			safeCount++
		}
	}

	return safeCount
}

func main() {
	file, err := os.Open("data")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reports := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		report := make([]int, len(parts))
		for i, part := range parts {
			report[i], _ = strconv.Atoi(part)
		}
		reports = append(reports, report)
	}

	result := countSafeReportsWithDampener(reports)
	fmt.Println("Safe reports:", result)
}
