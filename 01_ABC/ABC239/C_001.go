package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	d := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	if d == 20 || d == 2 || d == 16 || d == 4 || d == 18 || d == 10 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
