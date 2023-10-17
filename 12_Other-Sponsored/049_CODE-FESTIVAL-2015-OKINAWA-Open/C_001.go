package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	cnt := 0

	for i := 0; i < N; i++ {
		var s [3]string
		fmt.Scan(&s[0], &s[1], &s[2])
		b := [3]bool{}
		if i%2 == 0 {
			b[0] = s[0][0] == '#'
			b[1] = s[0][1] == '#'
			b[2] = s[0][2] == '#'
		} else {
			b[0] = s[0][0] == '#'
			b[1] = s[1][0] == '#'
			b[2] = s[2][0] == '#'
		}
		if b[0] && b[1] && b[2] {
			A[i] = 2
		} else if (!b[0] && b[1] && b[2]) || (b[0] && b[1] && !b[2]) {
			A[i] = 1
		} else {
			A[i] = 0
		}
		cnt += boolToInt(!b[0]) + boolToInt(!b[1]) + boolToInt(!b[2])
	}

	sum := 0
	for _, a := range A {
		sum ^= a
	}

	if (cnt%2 == 0) != (sum == 0) {
		fmt.Println("Snuke")
	} else {
		fmt.Println("Sothe")
	}

}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
