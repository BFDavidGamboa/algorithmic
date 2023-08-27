package diagonaldifference

import "fmt"

func diagonalDiffernec(arr [][]int32) (abs int32) {

	var a, b int32
	var lastn = arr[len(arr)-1]
	// var last = arr[][arr]
	for i, n := range arr {

		for j, m := range n {
			if i == j {
				a += m
			}

			if lastn[len(lastn)-1] == n[len(n)-1] && j == 0 {
				b += m
			}
		}

	}

	fmt.Println("a ", a)

	return abs
}
