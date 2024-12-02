package main
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Check line layout", line)
			continue
		}
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Check data: ", line)
			continue
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)
	similarityScore := 0
	startingIndex := 0
	for index := 0; index < len(list1); index++ {
		total := 0
		for i := startingIndex; i < len(list2); i++ {
			if(list1[index] == list2[i]) {
				total += 1
			}
			//Cool little optimization that lets us use a sliding window when doing the list comps
			if(list1[index] < list2[i]) {
				similarityScore = similarityScore + (total * list1[index])
				startingIndex = i
				break
			}
		}

    }
	fmt.Println("Total: ", similarityScore)
}