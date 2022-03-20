package main

import "fmt"

func main() {
	var N, X, Y int
	fmt.Scan(&N, &X, &Y)

	for a := 1; a < N+1; a++ {
		if Y%a == 0 {
			for b := 1; b < a+1; b++ {
				if Y/a%b == 0 {
					for c := 1; c < b+1; c++ {
						d := X - a - b - c
						if 1 <= d && d <= N && a*b*c*d == Y {
							fmt.Println("Yes")
							return
						}
					}
				}
			}
		}
	}
	fmt.Println("No")
}
