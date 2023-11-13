package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	S := make([]int, n)
	for i := range S {
		fmt.Fscan(in, &S[i])
	}
	sort.Ints(S)
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var m int
		fmt.Fscan(in, &m)
		ans := m + n
		for m > 0 {
			m--
			var v int
			fmt.Fscan(in, &v)
			tmp := 0
			if binarySearch(S, 0, len(S)-1, v) != -1 {
				tmp = 1
			}
			ans -= tmp
		}
		fmt.Fprintln(out, ans)
	}
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
