package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f func() int
	f = func() int {
		var S string
		fmt.Fscan(in, &S)
		y, _ := strconv.Atoi(S[0:4])
		m, _ := strconv.Atoi(S[5:7])
		d, _ := strconv.Atoi(S[8:])
		ret := 0
		for ny := 2022; ny < y; ny++ {
			ret += 365
			if ny%4 == 0 && (ny%100 != 0 || ny%400 == 0) {
				ret++
			}
		}
		for nm := 1; nm < m; nm++ {
			if nm == 2 {
				if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
					ret += 29
				} else {
					ret += 28
				}
			} else {
				ret += 31
				if nm == 4 || nm == 6 || nm == 9 || nm == 11 {
					ret--
				}
			}
		}
		return ret + d - 1
	}

	s := f()
	t := f()
	ans := 0
	for i := s; i <= t; i++ {
		if i%7 <= 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
