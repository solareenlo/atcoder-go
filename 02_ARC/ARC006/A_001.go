package main

import "fmt"

func main() {
	n := 6

	e := make([]int, n)
	for i := range e {
		fmt.Scan(&e[i])
	}

	var b int
	fmt.Scan(&b)

	l := make([]int, n)
	for i := range l {
		fmt.Scan(&l[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if e[i] == l[j] {
				cnt++
			}
		}
	}

	bonus := false
	for i := 0; i < n; i++ {
		if l[i] == b {
			bonus = true
		}
	}

	switch cnt {
	case 6:
		fmt.Println(1)
	case 5:
		if bonus {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	case 4:
		fmt.Println(4)
	case 3:
		fmt.Println(5)
	default:
		fmt.Println(0)
	}
}
