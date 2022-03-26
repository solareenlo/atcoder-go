package main

import (
	"bufio"
	"fmt"
	"os"
)

const A = 4000005

var (
	a  = [A/64 + 5]uint64{}
	mx int
)

func Count(x int) bool {
	if a[x/64]&(1<<(x&63)) != 0 {
		return true
	}
	return false
}

func move(x int) {
	u := mx / 64
	mx += x
	du := x / 64
	di := x & 63
	if di == 0 {
		for i := u; i >= 0; i-- {
			a[i+du] |= a[i]
		}
		return
	}
	for i := u; i >= 0; i-- {
		a[i+du+1] |= a[i] >> (64 - di)
		a[i+du] |= a[i] << di
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	cnt := make([]int, A)
	sum := 0
	for i := 1; i < n+1; i++ {
		var x int
		fmt.Fscan(in, &x)
		sum += x
		cnt[x]++
	}

	mx = 0
	a[0] = 1
	for i := 1; i < sum+1; i++ {
		for cnt[i] > 2 {
			cnt[i*2]++
			cnt[i] -= 2
		}
		for cnt[i] != 0 {
			move(i)
			cnt[i]--
		}
	}

	for i := (sum + 1) / 2; i < sum+1; i++ {
		if Count(i) {
			fmt.Println(i)
			break
		}
	}
}
