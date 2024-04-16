package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func separateInput(input string) ([]int, []string, error) {
	var nums []int
	var opers []string
	pattern := `(\d+)|([\+\-\*\/])`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		num, err := strconv.Atoi(match[1])
		if err == nil {
			nums = append(nums, num)
		} else if match[2] != "" {
			opers = append(opers, match[2])
		}
	}

	if len(nums) != len(opers)+1 {
		return nil, nil, fmt.Errorf("invalid expression: %s", input)
	}

	return nums, opers, nil
}

func calc(nums []int, opers []string) (int, error) {
	var finalResult int = int(nums[0])

	for i := range opers {
		switch opers[i] {
		case "+":
			finalResult += int(nums[i+1])
		case "-":
			finalResult -= int(nums[i+1])
		case "*":
			finalResult *= int(nums[i+1])
		case "/":
			if nums[i+1] == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			finalResult /= int(nums[i+1])
		default:
			return 0, fmt.Errorf("invalid operator: %s", opers[i])
		}
	}

	return finalResult, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input calculation (e.g. 1 + 2 / 3 * 4)")
	fmt.Print("> ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input", err)
		return
	}
	input = strings.TrimSpace(input) //remove trailing whitespace

	pattern := `^(\d+\s*[\+\-\*\/]\s*)+\d+$` //regexp (checks input)
	match, err := regexp.MatchString(pattern, input)
	if err != nil || !match {
		fmt.Println("invalid input")
		return
	}

	nums, opers, err := separateInput(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result, err := calc(nums, opers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result: %v \n", result)
}
