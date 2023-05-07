package main

import (
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	if h == 1 || w == 1 {
		fmt.Println(h * w)
	} else {
		hh := h/2 + h%2
		ww := w/2 + w%2
		fmt.Println(hh * ww)
	}
}
