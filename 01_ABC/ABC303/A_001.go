package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var x, y string
	fmt.Scan(&x, &y)
	for i := 0; i < n; i++ {
		if x[i] != y[i] {
			if !(x[i] == '1' && y[i] == 'l' || x[i] == '0' && y[i] == 'o' || x[i] == 'l' && y[i] == '1' || x[i] == 'o' && y[i] == '0') {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
