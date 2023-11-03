package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var a [1111]int
	a[0] = 1
	t := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		t ^= (1 << int(s[i]-'0'))
		ans += a[t]
		a[t]++
	}
	fmt.Println(ans)
}
