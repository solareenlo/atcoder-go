package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var k, n int
	fmt.Fscan(in, &k, &n)
	a := make([][]int, 51)
	for i := 0; i < n; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		q--
		a[p] = append(a[p], q)
	}

	for i := 0; i < k+1; i++ {
		sort.Slice(a[i], func(l, r int) bool {
			return a[i][l] < a[i][r]
		})
	}

	ans := (1 << (k + 1)) - 1
	for i := 0; i < k+1; i++ {
		for _, q := range a[i] {
			d := i
			x := q
			found := false
			for d > 0 {
				x /= 2
				d--
				if binarySearch(a[d], 0, len(a[d])-1, x) != -1 {
					found = true
					break
				}
			}
			if !found {
				ans -= (1 << (k - i + 1)) - 1
			}
		}
	}
	fmt.Println(ans)
}

func binarySearch(numbers []int, leftBound, rightBound, numberToFind int) int {
	if rightBound >= leftBound {
		midPoint := leftBound + (rightBound-leftBound)/2
		if numbers[midPoint] == numberToFind {
			return midPoint
		}
		if numbers[midPoint] > numberToFind {
			return binarySearch(numbers, leftBound, midPoint-1, numberToFind)
		}
		return binarySearch(numbers, midPoint+1, rightBound, numberToFind)
	}
	return -1
}
