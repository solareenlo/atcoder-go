package main

import (
	"fmt"
	"strings"
)

func main() {
	var w string
	fmt.Scan(&w)

	res := "Yes"
	for i := range w {
		if strings.Count(w, string(w[i]))%2 != 0 {
			res = "No"
		}
	}
	fmt.Println(res)
}
