package main

import "fmt"

func main() {
	var T1, T2, A1, A2, B1, B2 int
	fmt.Scan(&T1, &T2, &A1, &A2, &B1, &B2)

	A1 -= B1
	A2 -= B2

	B1 = T1 * A1
	B2 = T2 * A2
	if B1 < 0 {
		B1 *= -1
		B2 *= -1
	}

	if B1+B2 > 0 {
		fmt.Println(0)
	} else if B1+B2 == 0 {
		fmt.Println("infinity")
	} else if B1%(-B1-B2) != 0 {
		fmt.Println(B1/(-B1-B2)*2 + 1)
	} else {
		fmt.Println(B1 / (-B1 - B2) * 2)
	}
}
