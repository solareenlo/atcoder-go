package main

import "fmt"

func qry(x, y int) string {
	fmt.Println("?", x, y)
	var c string
	fmt.Scan(&c)
	return c
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a <= b {
		fmt.Println("Impossible")
		return
	}

	s := make([]int, 9999)
	n := a + b
	t := 0
	for i := 0; i < n; i++ {
		if t == 0 || qry(s[t], i) == "Y" {
			t++
			s[t] = i
		} else {
			t--
		}
	}

	v := make([]int, 9999)
	a = s[t]
	v[a] = 1
	for i := 0; i < n; i++ {
		if (i ^ a) != 0 {
			if qry(a, i) == "Y" {
				v[i] = 1
			} else {
				v[i] = 0
			}
		}
	}

	fmt.Print("! ")
	for i := 0; i < n; i++ {
		fmt.Print(v[i])
	}
}
