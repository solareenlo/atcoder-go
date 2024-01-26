package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	mp := make(map[string]int)
	mp["t"] = 3858
	mp["k"] = 3679
	mp["B"] = 3658
	mp["U"] = 3648
	mp["a"] = 3638
	mp["S"] = 3630
	mp["e"] = 3613
	mp["m"] = 3555
	mp["n"] = 3516
	mp["s"] = 3481
	fmt.Println(mp[string(s[0])])
}
