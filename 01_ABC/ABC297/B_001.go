package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var a [5][]int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'K':
			a[0] = append(a[0], i)
		case 'Q':
			a[1] = append(a[1], i)
		case 'R':
			a[2] = append(a[2], i)
		case 'B':
			a[3] = append(a[3], i)
		case 'N':
			a[4] = append(a[4], i)
		}
	}
	if len(a[0]) == 1 && len(a[1]) == 1 && len(a[2]) == 2 && len(a[3]) == 2 && len(a[4]) == 2 {
		if (a[3][0]+a[3][1])%2 == 0 {
			fmt.Println("No")
		} else if a[2][0] < a[0][0] && a[0][0] < a[2][1] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		fmt.Println("No")
	}
}
