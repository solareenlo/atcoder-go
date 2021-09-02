package main

import "fmt"

func main() {
	var l1, l2, l3 int
	fmt.Scan(&l1, &l2, &l3)
	fmt.Println(l1 ^ l2 ^ l3)
}
