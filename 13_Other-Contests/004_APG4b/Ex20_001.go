package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	p[0] = -1
	for i := 1; i < n; i++ {
		fmt.Scan(&p[i])
	}

	child := make([][]int, n)
	for i := 1; i < n; i++ {
		child[p[i]] = append(child[p[i]], i)
	}

	for i := 0; i < n; i++ {
		fmt.Println(cnt(child, i))
	}
}

func cnt(child [][]int, x int) int {
	if len(child[x]) == 0 {
		return 1
	}
	num := 1
	for _, c := range child[x] {
		num += cnt(child, c)
	}
	return num
}
