package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	res := make([]byte, 3)
	for i := 0; i < 3; i++ {
		res[i] = n[i]
		if n[i] == '1' {
			res[i] = '9'
		} else if n[i] == '9' {
			res[i] = '1'
		}
	}
	fmt.Println(string(res))
}
