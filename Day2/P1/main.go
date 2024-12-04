package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"	
    "strconv"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	file, err := os.Open("data")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
OuterLoop:
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		isIncreasing := true
		
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		if(num1 > num2) {
			isIncreasing = false
		}
		//fmt.Println("Analyzing: ", line)
		//fmt.Println("Increasing: ", isIncreasing)
		for i := 1; i < len(parts); i++ {
            num1, _ := strconv.Atoi(parts[i-1])
            num2, _ := strconv.Atoi(parts[i])
            if num1 > num2 && isIncreasing {
                fmt.Println("Should be increasing but isn't: ", line)
                count++
                continue OuterLoop
            }
            if num1 < num2 && !isIncreasing {
                fmt.Println("Should be decreasing but isn't: ", line)
                count++
                continue OuterLoop
            }
            dif := abs(num1 - num2)
            if !(dif <= 3 && dif >= 1) {
				fmt.Println("Difference is not 1, 2 or 3: ", line)
                count++
                continue OuterLoop
            }
        }
		
	}
	fmt.Println(1000-count)
}