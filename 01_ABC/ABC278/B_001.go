package main

import (
	"fmt"
)

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	for {
		if (h/10)*10+m/10 <= 23 && (h%10)*10+m%10 <= 59 {
			fmt.Println(h, m)
			return
		}
		m++
		if m == 60 {
			h++
			m = 0
		}
		if h == 24 {
			h = 0
		}
	}
}
