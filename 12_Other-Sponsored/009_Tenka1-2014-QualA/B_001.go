package main

import "fmt"

func main() {
	var t, a, d [1 << 20]int

	var s string
	fmt.Scan(&s)
	s += "----------"
	p := -5
	n := 5
	c := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if p+2 < i {
			if s[i] == 'N' && n > 0 {
				n--
				a[i+1]++
				d[i+1] += 10 + c/10
				t[i+6]++
			}
			if s[i] == 'C' && n > 2 {
				n -= 3
				p = i
				a[i+3]++
				d[i+3] += 50 + c/10*5
				t[i+8] += 3
			}
		}
		c += a[i]
		ans += d[i]
		n += t[i]
	}
	fmt.Println(ans)
}
