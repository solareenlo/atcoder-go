package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := a; i <= 919; i++ {
		if (i/100)*(i/10%10) == i%10 {
			fmt.Println(i)
			return
		}
	}
}
