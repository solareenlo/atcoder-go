package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	w := make([]string, n)
	for i := range w {
		fmt.Scan(&w[i])
	}

	word := map[string]struct{}{}
	for i := 0; i < n; i++ {
		word[w[i]] = struct{}{}
	}

	ok := true
	if len(word) != n {
		ok = false
	} else {
		for i := 0; i < n; i++ {
			if i == n-1 {
				continue
			}
			if w[i][len(w[i])-1] != w[i+1][0] {
				ok = false
			}
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
