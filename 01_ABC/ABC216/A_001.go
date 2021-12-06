package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	xy := strings.Split(s, ".")
	y, _ := strconv.Atoi(xy[1])

	switch {
	case 0 <= y && y <= 2:
		fmt.Print(xy[0], "-")
	case 3 <= y && y <= 6:
		fmt.Print(xy[0])
	default:
		fmt.Print(xy[0], "+")
	}
	fmt.Println()
}
