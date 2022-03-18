package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	v := make(map[byte]int)
	for i := range s {
		v[s[i]] = 1
	}

	if v['N']^v['S'] != 0 || v['W']^v['E'] != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
