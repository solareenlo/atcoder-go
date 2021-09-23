package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	res1, res2 := false, false
	if a[len(a)-1] == b[0] {
		res1 = true
	}
	if b[len(b)-1] == c[0] {
		res2 = true
	}

	if res1 && res2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
