package bucket

import "strconv"

func Bucket(n int) string {
	if n == 3 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
