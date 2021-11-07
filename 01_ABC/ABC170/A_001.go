package main

import "fmt"

func main() {
	res := 0
	for i := 0; i < 5; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if tmp == 0 {
			res = i + 1
		}
	}

	fmt.Println(res)
}
