package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n < 10 {
		fmt.Println(n + 9)
	} else {
		sum := 0
		n1 := n
		for n1 != 0 {
			sum += n1 % 10
			n1 /= 10
		}
		c := sum / 9
		suma := sum % 9
		for j := 0; j < c; j++ {
			suma = suma*10 + 9
		}
		if n != suma {
			fmt.Println(suma)
		} else {
			sumb := ((sum%9)+1)*10 + 8
			for j := 1; j < c; j++ {
				sumb = sumb*10 + 9
			}
			fmt.Println(sumb)
		}
	}
}
