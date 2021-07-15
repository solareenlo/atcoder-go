package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scan(&n)
	var k int
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		tmp, _ := strconv.Atoi(n)
		if tmp%200 == 0 {
			tmp /= 200
			n = strconv.Itoa(tmp)
		} else {
			n += "200"
		}
	}
	fmt.Println(n)
}
