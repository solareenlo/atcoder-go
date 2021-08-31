package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	res, _ := strconv.Atoi(a + b)
	fmt.Println(res * 2)
}
