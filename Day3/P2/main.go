package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(buffer []string) bool {
	//fmt.Println("Testing buffer: ", buffer)
	if(len(buffer) == 0) {
		return true
	}
	if(len(buffer) >= 1 && buffer[0] != "m") {
		return false
	}
	if(len(buffer) >= 2 && buffer[1] != "u") {
		return false
	}
	if(len(buffer) >= 3 && buffer[2] != "l") {
		return false
	}
	if(len(buffer) >= 4 && buffer[3] != "(") {
		return false
	}
	inSecondMode := false
	//fmt.Println("Doing number cases")
	for i := 4; i < len(buffer); i++ {
		if (buffer[i] == ","){
			inSecondMode = true
		}else if(inSecondMode && buffer[i] == ")") {
			return true
		}else if(!( buffer[i] >= "0" && buffer[i] <= "9")) {
			//fmt.Println("Returning false because of number case: ", buffer[i])
			return false
		}
	}
	//fmt.Println("Returning true because of base case")
	return true

}

func checkEnabler(buffer []string) bool {
	if len(buffer) == 0 {
		return false
	}
	str := strings.Join(buffer, "")
	return strings.HasPrefix("do()",str) || strings.HasPrefix("don't()",str)
		
}


func processBuffer(buffer []string) int {
	str := strings.Join(buffer[4:len(buffer)-1], "")
	parts := strings.Split(str, ",")
	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[1])
	return first * second
}

func processEnabler(buffer []string) bool {
	str := strings.Join(buffer, "")
	return str == "do()"
}

func main() {
	file, _ := os.Open("data")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mathBuffer := []string{}
	enablerBuffer := []string{}
	enable:=true
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			c := string(char)
			enablerBuffer = append(enablerBuffer, c)

			if(checkEnabler(enablerBuffer)) {
				if(c == ")") {
					fmt.Println("Enabler: ", enablerBuffer)
					enable = processEnabler(enablerBuffer)
					
				}
			}else{
				enablerBuffer = []string{}
				mathBuffer = append(mathBuffer, c)
				if(isValid(mathBuffer) && c == ")") {
					fmt.Println(mathBuffer)
					if(enable) {
						result += processBuffer(mathBuffer)
					}
					mathBuffer = []string{}
				}else if(!isValid(mathBuffer)) {
					mathBuffer = []string{}
				}
			}

			
		}
	}
	fmt.Println("Result: ", result)

}