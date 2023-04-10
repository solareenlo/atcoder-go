package main

import (
	"fmt"
	"reflect"
)

func main() {
	var S string
	fmt.Scan(&S)
	n := len(S)
	S = reverseString(S)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = int(S[i] - '0')
	}
	var res string
	cnt := 0
	for {
		if reflect.DeepEqual(a, make([]int, n)) {
			break
		}
		rs := 0
		for i := n - 1; i >= 0; i-- {
			a[i] += rs * 10
			rs = a[i] % 5
			a[i] /= 5
		}
		if rs != 0 {
			for i := 0; i < cnt; i++ {
				res += "5*"
			}
			res += string('0' + rs)
			res += "+"
		}
		cnt++
	}
	res = res[:len(res)-1]
	fmt.Println(res)
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
