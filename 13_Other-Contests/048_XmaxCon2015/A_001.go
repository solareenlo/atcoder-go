package main

import "fmt"

func main() {
	var a, b, t, N, X, C int
	fmt.Scan(&N, &X, &t, &a, &b, &C)

	AT := 1
	BT := 0

	for t > 0 {
		if t%2 != 0 {
			AT = (AT * a) % C
			BT = (a*BT + b) % C
		}
		b = (a*b + b) % C
		a = (a * a) % C
		t /= 2
	}

	ans := 0

	for i := 0; i < N; i++ {
		ans += X
		X = (AT*X + BT) % C
	}

	fmt.Println(ans)
}
