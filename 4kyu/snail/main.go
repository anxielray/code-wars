package main

import "fmt"

/* Given an n x n array, return the array elements arranged from outermost elements
to the middle element, traveling clockwise. */

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(snail(arr)) // Output: [1, 2, 3, 6, 9, 8, 7, 4, 5]
}

func snail(snailMap [][]int) []int {
	if len(snailMap) == 0 || len(snailMap[0]) == 0 {
		return []int{}
	}

	result := make([]int, 0, len(snailMap)*len(snailMap[0]))
	top, bottom := 0, len(snailMap)-1
	left, right := 0, len(snailMap[0])-1
	for top <= bottom && left <= right {
		// Traverse the top row
		for col := left; col <= right; col++ {
			result = append(result, snailMap[top][col])
		}
		top++

		// Traverse the rightmost column
		for row := top; row <= bottom; row++ {
			result = append(result, snailMap[row][right])
		}
		right--

		// Traverse the bottom row (if exists)
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, snailMap[bottom][col])
			}
			bottom--
		}

		// Traverse the leftmost column (if exists)
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, snailMap[row][left])
			}
			left++
		}
	}
	return result
}

/* Approach 2 */

/*
func Snail(snailMap [][]int) []int {

  xmin := 0
  ymin := 0
  xmax := len( snailMap[0]) - 1
  ymax := len( snailMap) - 1
  resultln := len(snailMap[0]) * len(snailMap)
  result := make([]int, 0)

  for len(result) < resultln {

    for x := xmin; x <= xmax; x++ {
      result = append(result, snailMap[ymin][x])
    }
    ymin++

    for y := ymin; y <= ymax; y++ {
      result = append(result, snailMap[y][xmax])
    }
    xmax--

    for x := xmax; x >= xmin; x-- {
      result = append(result, snailMap[ymax][x])
    }
    ymax--

    for y := ymax; y >= ymin; y-- {
      result = append(result, snailMap[y][xmin])
    }
    xmin++
  }
  return result
}

Approach 3

func Snail(snaipMap [][]int) []int {
	result := []int{}
	lines := len(snaipMap[0])*2 - 1
	moves := len(snaipMap[0])
	dx, dy := 1, 0
	x, y := -1, 0

	for ; lines > 0; lines-- {
		for range make([]int, moves) {
			x += dx
			y += dy
			result = append(result, snaipMap[y][x])
		}
		moves -= 1 & dx
		dx, dy = -dy, dx
	}

	return result
}
*/
