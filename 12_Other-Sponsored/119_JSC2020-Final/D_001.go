package main

import "fmt"

func main() {
	const mask = (1 << 20) - 1

	var n int
	fmt.Scan(&n)
	ans := make([]int, 0)
	Len := 20
	for p := 17; p >= 0; p-- {
		if ((n >> p) & 1) != 0 {
			ans = append(ans, (1<<Len)-1)
		} else {
			Len--
			ans = append(ans, 1<<Len)
		}
	}

	fmt.Println(len(ans))
	for _, e := range ans {
		fmt.Printf("%d ", e^mask)
	}
}
