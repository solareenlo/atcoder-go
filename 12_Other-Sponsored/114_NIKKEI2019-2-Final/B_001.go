package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := 0
	for i := 2; i < len(s); i++ {
		for j := i + 2; j+2 <= len(s); j += 2 {
			f := true
			for k := 0; k < (j-i)/2; k++ {
				if f && s[i+k] == s[i+(j-i)/2+k] {
					f = true
				} else {
					f = false
				}
			}
			if f {
				x := i - 1
				y := len(s) - 1
				for 0 < x && j < y {
					if s[x] != s[y] {
						break
					}
					a++
					x--
					y--
				}
			}
		}
	}
	fmt.Println(a)
}
