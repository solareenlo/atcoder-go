package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100007
	var str string
	fmt.Fscan(in, &str)
	n := len(str)
	str = " " + str

	flg := 0
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		a[i] = (a[i-1] + int(str[i]-'a') + 1) % 3
		if str[i] == str[i-1] {
			flg = 1
		}
	}

	if flg == 0 {
		fmt.Println(1)
		return
	}

	const mod = 1_000_000_007
	f := make([]int, N)
	las := make([]int, 3)
	for i := 1; i <= n; i++ {
		tmp := 0
		if a[i] >= 1 {
			tmp = 1
		}
		f[i] = (f[las[0]] + f[las[1]] + f[las[2]] + tmp - f[las[a[i]]] + mod) % mod
		las[a[i]] = i
	}

	fmt.Println(f[n])
}
