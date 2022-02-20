package main

import (
	"fmt"
)

func main() {
	var n, v int
	fmt.Scan(&n)
	e, o := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if i%2 != 0 {
			o[v]++
		} else {
			e[v]++
		}
	}

	keyE0, valE0 := maxElement(e)
	keyO0, valO0 := maxElement(o)

	res := 0
	if keyE0 != keyO0 {
		res = n - valE0 - valO0
	} else {
		e[keyE0], o[keyO0] = 0, 0
		_, valE1 := maxElement(e)
		_, valO1 := maxElement(o)
		res = min(n-valE0-valE1, n-valO0-valO1)
	}
	fmt.Println(res)
}

func maxElement(m map[int]int) (int, int) {
	key, val := 0, 0
	for k, v := range m {
		if val < v {
			val, key = v, k
		}
	}
	return key, val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
