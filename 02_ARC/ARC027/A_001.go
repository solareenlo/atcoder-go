package main

import "fmt"

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	res := 0
	res += (17 - h) * 60
	res += 60 - m

	fmt.Println(res)
}
