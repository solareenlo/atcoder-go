package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := make([]int, 4)
	for i := range a {
		a[i] = int(s[i] - '0')
	}

	op := make([]string, 3)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				if i == 0 {
					a[1] *= -1
				}
				if j == 0 {
					a[2] *= -1
				}
				if k == 0 {
					a[3] *= -1
				}
				if a[0]+a[1]+a[2]+a[3] == 7 {
					if i == 0 {
						op[0] = "-"
					} else {
						op[0] = "+"
					}
					if j == 0 {
						op[1] = "-"
					} else {
						op[1] = "+"
					}
					if k == 0 {
						op[2] = "-"
					} else {
						op[2] = "+"
					}
				}
				for i := range a {
					a[i] = int(s[i] - '0')
				}
			}
		}
	}

	fmt.Print(a[0], op[0], a[1], op[1], a[2], op[2], a[3], "=7\n")
}
