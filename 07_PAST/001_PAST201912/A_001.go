package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		fmt.Println(i * 2)
	} else {
		fmt.Println("error")
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
