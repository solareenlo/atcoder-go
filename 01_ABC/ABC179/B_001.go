package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	ok := false
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		if l == r {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == 3 {
			ok = true
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
