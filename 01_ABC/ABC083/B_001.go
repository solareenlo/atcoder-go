package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	res := 0
	for i := 0; i <= n; i++ {
		tmp := strconv.Itoa(i)
		sum := 0
		for j := 0; j < len(tmp); j++ {
			sum += int(tmp[j] - '0')
		}
		if a <= sum && sum <= b {
			res += i
		}
	}
	fmt.Println(res)
}
