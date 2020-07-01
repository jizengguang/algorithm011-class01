package Week_02

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, 0)
	for i := 1; i <= n; i++ {
		var tmp = ""
		var three = i%3 == 0
		var five = i%5 == 0

		if three {
			tmp += "Fizz"
		}

		if five {
			tmp += "Buzz"
		}

		if tmp == "" {
			tmp = strconv.Itoa(i)
		}

		result = append(result, tmp)
	}
	return result
}
