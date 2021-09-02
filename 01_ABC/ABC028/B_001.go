package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := make([]int, 6)
	for i := range s {
		m[s[i]-'A']++
	}
	res := fmt.Sprint(m)
	fmt.Println(res[1 : len(res)-1])
}
