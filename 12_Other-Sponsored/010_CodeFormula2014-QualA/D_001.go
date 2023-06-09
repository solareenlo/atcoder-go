package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, k string
	fmt.Scan(&s, &k)
	cnt := 0
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s[0:i], s[i]) == -1 && strings.IndexByte(k, s[i]) == -1 {
			cnt++
		}
	}
	t := 1.0
	ans := 0.0
	for i := 1; i < 36-len(k); i++ {
		var tmp float64
		if i > cnt {
			tmp = 1
		} else {
			tmp = 1.0 - 1.0/(float64(i)+1)
		}
		t *= tmp
		ans += 1 - t
	}
	fmt.Println(ans*2 + float64(len(s)))
}
