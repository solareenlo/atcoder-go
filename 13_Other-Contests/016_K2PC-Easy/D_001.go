package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var K, N int
	fmt.Fscan(in, &K, &N)
	t := make([]int, N)
	p := make([]int, N)
	q := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &p[i], &q[i])
		q[i]--
		t = append(t, (1<<p[i])+q[i])
	}
	sort.Ints(t)
	ans := (1 << (K + 1)) - 1
	for i := 0; i < N; i++ {
		cur := (1 << p[i]) + q[i]
		f := true
		if cur != 1 {
			for {
				cur >>= 1
				if binarySearch(t, 0, len(t), cur) != -1 {
					f = false
					break
				}
				if cur == 1 {
					break
				}
			}
		}
		if f {
			ans -= (1 << (K - p[i] + 1)) - 1
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
