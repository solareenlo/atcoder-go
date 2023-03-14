package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b int
	fmt.Fscan(in, &a, &b)

	ret := 0
	for i := 0; i < b; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		p--
		q--
		T := 0
		tmp := a - 1
		for tmp > 0 {
			T += tmp / 2
			tmp /= 2
		}
		tmp = a - 1 - p
		for tmp > 0 {
			T -= tmp / 2
			tmp /= 2
		}
		tmp = q
		for tmp > 0 {
			T -= tmp / 2
			tmp /= 2
		}
		tmp = p - q
		for tmp > 0 {
			T -= tmp / 2
			tmp /= 2
		}
		if T == 0 {
			if ret == 0 {
				ret = 1
			} else {
				ret = 0
			}
		}
	}

	if a%2 == 0 {
		if ret == 0 {
			ret = 1
		} else {
			ret = 0
		}
	}

	if ret != 0 {
		fmt.Println("Iori")
	} else {
		fmt.Println("Yayoi")
	}
}
