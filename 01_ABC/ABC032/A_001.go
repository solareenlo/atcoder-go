package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	var res int
	for i := 0; i < 20001; i++ {
		res = n + i
		if res%a == 0 && res%b == 0 {
			fmt.Println(res)
			break
		}
	}
}
