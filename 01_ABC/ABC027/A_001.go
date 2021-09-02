package main

import "fmt"

func main() {
	l := make([]int, 3)
	for i := range l {
		fmt.Scan(&l[i])
	}
	switch l[0] {
	case l[1]:
		fmt.Println(l[2])
	case l[2]:
		fmt.Println(l[1])
	default:
		fmt.Println(l[0])
	}
}
