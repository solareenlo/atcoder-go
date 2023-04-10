package main

import (
	"bufio"
	"fmt"
	"os"
)

var kanji []string = []string{"〇", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "百", "千", "万", "億", "の", "乗"}
var N int
var s []int
var mphi map[int]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x string
	fmt.Fscan(in, &x)
	N = len(x) / 3
	for i := 0; i < N; i++ {
		c := x[i*3 : i*3+3]
		for j := 0; j < len(kanji); j++ {
			if c == kanji[j] {
				s = append(s, j)
			}
		}
	}
	N++
	s = append(s, 100)

	it := 0
	MOD := int(1e9 + 9)
	mphi = make(map[int]int)
	v := eval(&it, MOD) % MOD

	ans := make([]int, 0)

	var easy func(int)
	easy = func(x int) {
		q := x / 1000
		if q != 0 {
			if q >= 2 {
				ans = append(ans, q)
			}
			ans = append(ans, 12)
		}
		x %= 1000
		q = x / 100
		if q != 0 {
			if q >= 2 {
				ans = append(ans, q)
			}
			ans = append(ans, 11)
		}
		x %= 100
		q = x / 10
		if q != 0 {
			if q >= 2 {
				ans = append(ans, q)
			}
			ans = append(ans, 10)
		}
		x %= 10
		q = x
		if q != 0 {
			ans = append(ans, q)
		}
	}

	if v == 0 {
		ans = append(ans, 0)
	} else {
		q := v / 100000000
		if q != 0 {
			easy(q)
			ans = append(ans, 14)
		}
		v %= 100000000
		q = v / 10000
		if q != 0 {
			easy(q)
			ans = append(ans, 13)
		}
		v %= 10000
		q = v
		if q != 0 {
			easy(q)
		}
	}
	for _, i := range ans {
		fmt.Fprint(out, kanji[i])
	}
	fmt.Fprintln(out)
}

func eval(it *int, MOD int) int {
	x := num(it, MOD)
	for s[*it] == 15 {
		*it++
		y := eval(it, phi(MOD))
		x = ex(x, y, MOD)
		*it++
	}
	return x
}

func num(it *int, MOD int) int {
	res := 0
	for s[*it] < 15 {
		x := num2(it)
		if s[*it] == 13 {
			x *= 10000
			*it++
		}
		if s[*it] == 14 {
			x *= 10000 * 10000
			*it++
		}
		res += x
	}
	return take_mod(res, MOD)
}

func num2(it *int) int {
	res := 0
	for s[*it] <= 12 {
		x := 1
		if s[*it] < 10 {
			x = s[*it]
			*it++
		}
		if s[*it] == 10 {
			x *= 10
			*it++
		}
		if s[*it] == 11 {
			x *= 100
			*it++
		}
		if s[*it] == 12 {
			x *= 1000
			*it++
		}
		res += x
	}
	return res
}

func phi(x int) int {
	if _, ok := mphi[x]; ok {
		return mphi[x]
	}
	_x := x
	phix := x
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			for x%i == 0 {
				x /= i
			}
			phix /= i
			phix *= i - 1
		}
	}
	if x > 1 {
		phix /= x
		phix *= x - 1
	}
	mphi[_x] = phix
	return mphi[_x]
}

func ex(x, y, MOD int) int {
	a := 1
	for y != 0 {
		if (y & 1) != 0 {
			a = take_mod(a*x, MOD)
		}
		x = take_mod(x*x, MOD)
		y /= 2
	}
	return a
}

func take_mod(x, MOD int) int {
	if x < 100 || x < MOD {
		return x
	}
	return x%MOD + MOD
}
