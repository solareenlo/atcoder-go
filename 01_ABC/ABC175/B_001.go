package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	l := make([]int, n)
	for i := range l {
		fmt.Scan(&l[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if l[i] == l[j] || l[j] == l[k] || l[k] == l[i] {
					continue
				}
				if !(l[i]+l[j] <= l[k] || l[i]+l[k] <= l[j] || l[j]+l[k] <= l[i]) {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}
