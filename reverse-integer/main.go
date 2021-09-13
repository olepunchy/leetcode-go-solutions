package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	// Typical number
	number := 123
	result := reverse(number)
	fmt.Println(result)

	// Large positive number, should return 0 7463847421 is larger than Int32
	anotherNumber := 2147483647
	anotherResult := reverse(anotherNumber)
	fmt.Println(anotherResult)

	// Big negative number, should return 0 -8463847412 is larger than Int32
	bigNegative := -2147483648
	bigNegativeReversed := reverse(bigNegative)
	fmt.Println(bigNegativeReversed)
}

func numberToString(number int) string {
	return strconv.Itoa(number)
}

func reverseString(input string) string {

	letters := []rune(input)
	for i, j := 0, len(letters) - 1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	return string(letters)
}

func reverseNegativeSign(input string) string {
	trimmedNumberString := strings.TrimSuffix(
		input,
		"-",
		)

	stringBuilder := new(strings.Builder)
	stringBuilder.WriteString("-")
	stringBuilder.WriteString(trimmedNumberString)

	return stringBuilder.String()
}

func reverse(number int) int {
	if number == 0 {
		return 0
	}

	positive := number > 0

	numberString := numberToString(number)
	reversedNumber := reverseString(numberString)

	if !positive {
		resultString := reverseNegativeSign(reversedNumber)

		result, _ := strconv.Atoi(resultString)

		if result < -math.MaxInt32 {
			return 0
		} else  {
			return result
		}
	} else {
		result, _ := strconv.Atoi(reversedNumber)

		if result > math.MaxInt32 {
			return 0
		} else {
			return result
		}
	}

}