package main

import "fmt"

func out(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(s)
	}
}

func main() {
	var sx, sy, tx, ty int
	fmt.Scan(&sx, &sy, &tx, &ty)

	dy, dx := ty-sy, tx-sx
	out("U", dy)
	out("R", dx)
	out("D", dy)
	out("L", dx)
	out("L", 1)
	out("U", dy+1)
	out("R", dx+1)
	out("D", 1)
	out("R", 1)
	out("D", dy+1)
	out("L", dx+1)
	out("U", 1)
	out("\n", 1)
}
