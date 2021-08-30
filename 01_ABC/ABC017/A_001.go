package main

import "fmt"

func main() {
	var s [3]int = [3]int{}
	var e [3]int = [3]int{}
	for i := 0; i < 3; i++ {
		fmt.Scan(&s[i], &e[i])
	}
	fmt.Println(float64(s[0]*e[0]+s[1]*e[1]+s[2]*e[2]) * 0.1)
}
