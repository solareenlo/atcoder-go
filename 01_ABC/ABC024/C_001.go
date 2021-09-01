package main

import "fmt"

func main() {
	var n, d, k int
	fmt.Scan(&n, &d, &k)

	l := make([]int, d)
	r := make([]int, d)
	for i := 0; i < d; i++ {
		fmt.Scan(&l[i], &r[i])
	}
	s := make([]int, k)
	t := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&s[i], &t[i])
	}

	for i := 0; i < k; i++ {
		res := -1
		for j := 0; j < d; j++ {
			if s[i] >= l[j] && s[i] <= r[j] {
				if t[i] >= l[j] && t[i] <= r[j] {
					s[i] = t[i]
				} else {
					if s[i] < t[i] {
						s[i] = r[j]
					} else {
						s[i] = l[j]
					}
				}
			}
			if s[i] == t[i] {
				res = j + 1
				break
			}
		}
		fmt.Println(res)
	}
}
