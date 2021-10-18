package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([][]int, m)
	for i := 0; i < m; i++ {
		var k int
		fmt.Scan(&k)
		s[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&s[i][j])
			s[i][j]--
		}
	}

	p := make([]int, m)
	for i := range p {
		fmt.Scan(&p[i])
	}

	res := 0
	for bit := 0; bit < 1<<n; bit++ {
		ok := true
		for j := 0; j < m; j++ {
			cnt := 0
			for _, id := range s[j] {
				if (bit>>id)&1 == 1 {
					cnt++
				}
			}
			cnt %= 2
			if cnt != p[j] {
				ok = false
			}
		}
		if ok {
			res++
		}
	}
	fmt.Println(res)
}
