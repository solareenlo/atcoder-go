package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	cnt := 0
	for y := 1; y <= r+1; y++ {
		for x := 1; x <= r+1; x++ {
			if (x*c)*(x*c)+(y*c)*(y*c) <= r*r {
				cnt++
			}
		}
	}
	fmt.Println(cnt << 2)
}
