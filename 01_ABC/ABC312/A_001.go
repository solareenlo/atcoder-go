package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "ACEBDFCEGDFAEGBFACGBD"
	var b string
	fmt.Scan(&b)
	if strings.Index(a, b) == -1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
