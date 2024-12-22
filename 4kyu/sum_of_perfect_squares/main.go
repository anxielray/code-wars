package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/anxielray/packages/IntPow"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:go run . <number>")
		os.Exit(1)
	}

	n, _ := strconv.Atoi(strings.TrimSpace(os.Args[1]))
	fmt.Println(sum_of_squares(n))
}

func sum_of_squares(n int) int {
	var result int

	for i := 1; i <= n; i++ {
		if isPerfectSquare(i) {
			result += i * i
		}
	}

	return result
}

func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	sqrt := IntPow.Sqrt(n)
	return sqrt*sqrt == n
}
