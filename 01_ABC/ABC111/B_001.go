package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if n[0] == n[1] && n[1] == n[2] {
		fmt.Println(n)
	} else {
		tmp := make([]byte, 0)
		for i := 0; i < 3; i++ {
			tmp = append(tmp, n[0])
		}
		if string(tmp) > n {
			fmt.Println(string(tmp))
		} else {
			res := make([]byte, 0)
			for i := 0; i < 3; i++ {
				res = append(res, n[0]+1)
			}
			fmt.Println(string(res))
		}
	}
}
