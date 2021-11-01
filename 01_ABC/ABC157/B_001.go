package main

import "fmt"

func main() {
	a := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	var n int
	fmt.Scan(&n)

	for k := 0; k < n; k++ {
		var b int
		fmt.Scan(&b)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if a[i][j] == b {
					a[i][j] = 0
				}
			}
		}
	}

	ok := false
	for i := 0; i < 3; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			sum += a[i][j]
		}
		if sum == 0 {
			ok = true
		}
	}

	for j := 0; j < 3; j++ {
		sum := 0
		for i := 0; i < 3; i++ {
			sum += a[i][j]
		}
		if sum == 0 {
			ok = true
		}
	}

	if a[0][0]+a[1][1]+a[2][2] == 0 {
		ok = true
	}
	if a[0][2]+a[1][1]+a[2][0] == 0 {
		ok = true
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
