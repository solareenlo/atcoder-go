package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	str := "atcoder"
	fmt.Println(str[l-1 : r])
}
