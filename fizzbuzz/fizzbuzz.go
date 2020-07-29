package fizzbuzz

import (
	"fmt"
	"strconv"
)

func Fizzbuzz(i int) string {
	return strconv.Itoa(i)
}

func FizzBuzzHackerrank(n int32) {
	for i := int32(1); i <= n; i++ {
		var result string
		if i%3 == 0 {
			result += "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}
		if result == "" {
			result = fmt.Sprintf("%d", i)
		}

		fmt.Println(result)
	}
}
