package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	X := 1
	a := [3]int{1, 2, 3}
	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		var S string
		fmt.Fscan(in, &S)
		x, y, z := a[0], a[1], a[2]
		if S == "East" {
			a[0] = 7 - z
			a[2] = x
		}
		if S == "Left" {
			a[1] = 7 - z
			a[2] = y
		}
		if S == "North" {
			a[0] = y
			a[1] = 7 - x
		}
		if S == "Right" {
			a[1] = z
			a[2] = 7 - y
		}
		if S == "South" {
			a[0] = 7 - y
			a[1] = x
		}
		if S == "West" {
			a[0] = z
			a[2] = 7 - x
		}
		X += a[0]
	}
	fmt.Println(X)
}
