package main

import "fmt"

func main() {
	var n, ng int
	fmt.Scan(&n)

	set := map[int]struct{}{}
	for i := 0; i < 3; i++ {
		fmt.Scan(&ng)
		set[ng] = struct{}{}
	}

	if _, ok := set[n]; ok {
		fmt.Println("NO")
		return
	}

	for i := 100; i > 0; i-- {
		n -= 3
		for {
			if _, ok := set[n]; ok {
				n++
			} else {
				break
			}
		}
	}

	if n <= 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
