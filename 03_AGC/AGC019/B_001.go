package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	dem := map[byte]int{}
	d := 0
	for i := 0; i < len(s); i++ {
		d += (i - dem[s[i]])
		dem[s[i]]++
	}

	fmt.Println(d + 1)
}
