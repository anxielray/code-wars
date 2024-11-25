package main

import (
	"fmt"
)

/* In number theory and combinatorics, a partition of a positive integer n, also called an integer partition,
is a way of writing n as a sum of positive integers. Two sums that differ only in the order of their summands
are considered the same partition. If order matters, the sum becomes a composition. For example, 4 can be partitioned
in five distinct ways: */

func partitionCount(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	// Create a DP array to store partition counts for each number up to n
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			dp[j] += dp[j-i]
		}
	}

	return dp[n]
}

func main() {
	n := 4
	fmt.Printf("The number of partitions for %d is: %d\n", n, partitionCount(n))
}
