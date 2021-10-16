package main

import "fmt"

func main() {
	var b string
	fmt.Scan(&b)

	m := make(map[string]string)
	m["A"] = "T"
	m["T"] = "A"
	m["C"] = "G"
	m["G"] = "C"

	fmt.Println(m[b])
}
