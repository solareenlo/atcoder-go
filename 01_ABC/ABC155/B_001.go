package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cntE := 0
	cnt := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a%2 == 0 {
			cntE++
			if a%3 == 0 || a%5 == 0 {
				cnt++
			}
		}
	}

	if cntE == cnt {
		fmt.Println("APPROVED")
	} else {
		fmt.Println("DENIED")
	}
}
