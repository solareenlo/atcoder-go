package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, m int
	fmt.Fscan(in, &n, &x, &m)
	l := make([]int, m)
	r := make([]int, m)
	s := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l[i], &r[i], &s[i])
		l[i]--
	}
	mx := -1
	ans := 0
	for S := 0; S < powMod(x+1, n); S++ {
		S2 := S
		var arr [6]int
		for i := 0; i < n; i++ {
			arr[i] = S2 % (x + 1)
			S2 /= x + 1
		}
		ok := true
		for i := 0; i < m; i++ {
			sum := 0
			for j := l[i]; j < r[i]; j++ {
				sum += arr[j]
			}
			if sum != s[i] {
				ok = false
			}
		}
		if ok {
			sum := 0
			for i := 0; i < n; i++ {
				sum += arr[i]
			}
			if sum > mx {
				mx = sum
				ans = S
			}
		}
	}
	if mx == -1 {
		fmt.Print(-1)
		return
	}
	arr := make([]int, 6)
	for i := 0; i < n; i++ {
		arr[i] = ans % (x + 1)
		ans /= x + 1
	}
	for i := 0; i < n; i++ {
		if i+1 == n {
			fmt.Println(arr[i])
		} else {
			fmt.Print(arr[i], " ")
		}
	}
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
