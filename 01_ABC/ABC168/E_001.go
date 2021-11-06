package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	zero := 0

	type pair struct{ x, y int }
	mp := map[pair]pair{}
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if x == 0 && y == 0 {
			zero++
			continue
		}
		g := gcd(x, y)
		x /= g
		y /= g
		if y < 0 {
			x, y = -x, -y
		}
		if y == 0 && x < 0 {
			x, y = -x, -y
		}
		rot90 := false
		if x <= 0 {
			rot90 = true
		}
		if rot90 {
			x, y = y, -x
		}
		p := mp[pair{x, y}]
		if rot90 {
			p.x++
		} else {
			p.y++
		}
		mp[pair{x, y}] = p
	}

	res := 1
	for _, p := range mp {
		s, t := p.x, p.y
		now := 1 // s, t どちらからも選ばない個数
		now += powMod(2, s) - 1
		now += mod
		now %= mod
		now += powMod(2, t) - 1
		now += mod
		now %= mod
		res *= now
		res %= mod
	}

	res -= 1 // 誰も選ばない通りを引いた
	res += mod
	res %= mod
	res += zero // 0単体の場合を足す
	res %= mod

	fmt.Println(res)
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
