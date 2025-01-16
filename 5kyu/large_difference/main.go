/*
Summary: Write a function which takes an array data of numbers and returns
the largest difference in indexes j - i such that data[i] <= data[j]

Long Description:

The largestDifference takes an array of numbers. That array is not sorted. Do not sort it
or change the order of the elements in any way, or their values.

Consider all of the pairs of numbers in the array where the first one is less than or equal
to the second one.

From these, find a pair where their positions in the array are farthest apart.

Return the difference between the indexes of the two array elements in this pair.

Example:

[ 1, 2, 3] returns 2 because here j = 2 and i = 0 and 2 - 0 = 2

*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 4, 1, 10, 3, 4, 0, -1, -2}
	fmt.Println(LargestDifference(arr)) // Output: 7
}

func LargestDifference(data []int) int {

	if len(data) == 0 {
		return 0
	} else if len(data) == 1 {
		return 0
	} else if len(data) == 2 {
		return 1
	} else if len(data) > 2 {
		return filter2(filter1(data), data)
	}
	return 0
}

//=============OUTLINE=============
//filter the elements with a proceeding element which is not less that the current element.
func filter1(a []int) [][]int {
	var wanted []int
	var finals [][]int
	for j := 0; j < len(a); j++ {

		for i := j + 1; i < len(a); i++ {
			if a[j] <= a[i] {
				wanted = append(wanted, a[j])
				wanted = append(wanted, a[i])
				finals = append(finals, wanted)
				wanted = nil
			}
		}
	}
	return finals
}

//further filter the array by finding the larges gap difference

func filter2(a [][]int, original []int) int {
	var (
		greater int
	)

	for _, r := range a {

		if r[0] == r[1] {
			if findNextindex(r[1], findIndex(r[0], original), original)-findIndex(r[0], original) > greater {
				greater = findNextindex(r[1], findIndex(r[0], original), original)-findIndex(r[0], original)
			}
		}else {

			if findIndex(r[1], original)-findIndex(r[0], original) > greater {
				greater = findIndex(r[1], original) - findIndex(r[0], original)
			}
		}

	}
	return greater
}

func findIndex(n int, a []int) int {
	for i, v := range a {
		if v == n {
			return i
		}
	}
	return -1
}

func findNextindex(n, from int, a[]int) int{
	for i := from + 1; i < len(a); i++ {
        if a[i] == n {
            return i
        }
    }
    return -1
}
