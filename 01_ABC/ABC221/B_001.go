package main

import (
	"fmt"
)

func main() {
	s := make([]byte, 0)
	t := make([]byte, 0)
	fmt.Scan(&s, &t)

	if equalSlice(s, t) {
		fmt.Println("Yes")
		return
	}

	n := len(s)
	for i := 0; i < n-1; i++ {
		tmp := swapSlice(s, i, i+1)
		if equalSlice(tmp, t) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func equalSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func swapSlice(bs []byte, i, j int) []byte {
	res := make([]byte, len(bs))
	copy(res, bs)
	res[i], res[j] = res[j], res[i]
	return res
}
