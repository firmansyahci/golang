package main

import (
	"fmt"
	"strconv"
)

func main() {
	var start int64
	var end int64
	count := 0

	fmt.Print("input start end number : ")
	fmt.Scanf("%d %d", &start, &end)

	for i := start; i <= end; i++ {
		x := strconv.FormatInt(i, 10)
		lastIndex := len(strconv.FormatInt(i, 10)) - 1
		isPalindrom := true

		for j := 0; j < len(strconv.FormatInt(i, 10)); j++ {
			if string(x[j]) != string(x[lastIndex]) {
				isPalindrom = false
				break
			}

			lastIndex--
		}

		if isPalindrom {
			count++
		}
	}

	fmt.Print(count)
}
