package main

import "fmt"

func main() {
	var P int
	fmt.Scan(&P)

	A, B, cnt := 1, 1, 1
	for A%P != 0 {
		tmp := A
		A = B
		B = (B + tmp) % P
		cnt++
	}
	fmt.Println(cnt)
}
