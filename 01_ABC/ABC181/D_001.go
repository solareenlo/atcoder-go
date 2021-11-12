package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	ok := false

	if n == 1 {
		if s[0] == '8' {
			ok = true
		}
	} else if n == 2 {
		t, _ := strconv.Atoi(s)
		if t%8 == 0 {
			ok = true
		}
		a, b := string(s[0]), string(s[1])
		s = b + a + s[2:]
		t, _ = strconv.Atoi(s)
		if t%8 == 0 {
			ok = true
		}
	} else {
		cnt := make([]int, 10)
		for _, x := range s {
			cnt[x-'0']++
		}
		for i := 112; i < 1000; i += 8 {
			c := make([]int, 10)
			copy(c, cnt)
			for _, x := range strconv.Itoa(i) {
				c[x-'0']--
			}
			tmp := 0
			for j := 0; j < 10; j++ {
				if c[j] >= 0 {
					tmp++
				}
			}
			if tmp == 10 {
				ok = true
			}
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
