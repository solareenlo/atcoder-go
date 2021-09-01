package main

import "fmt"

func main() {
	var n, a, b, k, p int
	fmt.Scan(&n, &a, &b, &k)

	s := map[int]struct{}{}
	s[a] = struct{}{}
	s[b] = struct{}{}
	for i := 0; i < k; i++ {
		fmt.Scan(&p)
		s[p] = struct{}{}
	}

	if len(s) == k+2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
