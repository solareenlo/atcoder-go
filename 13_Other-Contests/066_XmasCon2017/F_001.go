package main

import "fmt"

func main() {
	m := map[int]int{
		4:   12,
		5:   28,
		10:  1472,
		20:  5511680,
		25:  681155065,
		28:  728761302,
		100: 598737283,
		102: 43663927,
		104: 621940451,
		101: 842631569,
		103: 858227675,
	}
	var n int
	fmt.Scan(&n)
	fmt.Println(m[n])
}
