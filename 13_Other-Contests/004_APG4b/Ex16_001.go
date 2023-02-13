package main

import "fmt"

func main() {
	var data [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&data[i])
	}

	for i := 0; i < 4; i++ {
		if data[i] == data[i+1] {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
