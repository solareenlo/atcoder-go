package main

import (
	"fmt"
	"strconv"
)

func main() {
	var m int
	fmt.Scan(&m)

	vv := ""
	switch {
	case m < 100:
		vv = "00"
	case m < 1000:
		vv += "0"
		vv += strconv.Itoa(m / 100)
	case m <= 5000:
		vv = strconv.Itoa(m / 100)
	case m <= 30000:
		vv = strconv.Itoa(m/1000 + 50)
	case m <= 70000:
		vv = strconv.Itoa((m/1000-30)/5 + 80)
	default:
		vv = "89"
	}
	fmt.Println(vv)
}
