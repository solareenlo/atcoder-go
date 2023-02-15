package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Scan(&a)
	if a == 2 {
		fmt.Println(0)
		return
	}
	if math.Sqrt(a) != float64(int(math.Sqrt(a))) {
		fmt.Println(-1)
		return
	}
	now := 2
	ans := 0
	for s := int(math.Sqrt(a)); s >= 2; s-- {
		ans += s*s - now + 1
		now = s
	}
	fmt.Println(ans)
}
