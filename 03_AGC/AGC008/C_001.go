package main

import "fmt"

func main() {
	var I, O, T, J, L, S, Z int
	fmt.Scan(&I, &O, &T, &J, &L, &S, &Z)

	x := (I%2 + J%2 + L%2) % 3
	if I != 0 && J != 0 && L != 0 && x == 2 {
		x = 1
	}
	fmt.Println(I + J + L + O - x)
}
