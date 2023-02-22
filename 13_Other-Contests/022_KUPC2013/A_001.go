package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s [105]string
	var a [105]int
	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &s[i])
	}
	if q < a[1] {
		fmt.Println("kogakubu10gokan")
		return
	}
	if q >= a[n] {
		fmt.Println(s[n])
		return
	}
	for i := 1; i < n; i++ {
		if q >= a[i] && q < a[i+1] {
			fmt.Println(s[i])
		}
	}
}
