package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	diff := abs(b - a)
	div := diff / 10
	rem := diff % 10

	ans := div
	switch rem {
	case 1:
		ans += 1
	case 2:
		ans += 2
	case 3:
		ans += 3
	case 4:
		ans += 2
	case 5:
		ans += 1
	case 6:
		ans += 2
	case 7:
		ans += 3
	case 8:
		ans += 3
	case 9:
		ans += 2
	}

	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
