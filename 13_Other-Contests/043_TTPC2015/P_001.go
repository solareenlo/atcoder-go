package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAX_N = 5000

var N int
var iter int
var S string

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &S)
	S += " "

	iter = 0
	result := parse()
	sort.Slice(result, func(i, j int) bool {
		if result[i].a == result[j].a {
			return result[i].b < result[j].b
		}
		return result[i].a < result[j].a
	})
	result = reverseOrderPair(result)

	mx_a := 0
	if len(result) != 0 {
		mx_a = result[0].a
	}
	if mx_a == 0 {
		fmt.Println("YES")
		return
	}

	table := make([][]bool, MAX_N)
	for i := range table {
		table[i] = make([]bool, MAX_N)
	}
	sz := len(result)
	before_a := mx_a
	for i := 0; i < sz; i++ {
		p := result[i]
		if before_a != p.a {
			minimize(before_a, table)
		}
		table[p.a][p.b] = true
		before_a = p.a
	}
	minimize(before_a, table)
	if table[1][0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func parse() []prix {
	c := S[iter]
	ret := make([]prix, 0)
	if c == '(' {
		iter++
		left := parse()
		o := S[iter]
		iter++
		right := parse()
		iter++
		if o == ';' {
			ret = semicolon_merge(left, right)
		}
		if o == '+' {
			ret = plus_merge(left, right)
		}
	} else if c == '*' {
		iter++
		iter++
		right := parse()
		iter++
		ret = asterisk_merge(right)
	} else if c == 'a' {
		cnt := 0
		for S[iter] == 'a' && iter < N {
			cnt++
			iter++
		}
		ret = append(ret, prix{0, cnt})
	}
	sort.Slice(ret, func(i, j int) bool {
		if ret[i].a == ret[j].a {
			return ret[i].b < ret[j].b
		}
		return ret[i].a < ret[j].a
	})
	ret = uniquePair(ret)
	return ret
}

func uniquePair(a []prix) []prix {
	occurred := map[prix]bool{}
	result := []prix{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].a == result[j].a {
			return result[i].b < result[j].b
		}
		return result[i].a < result[j].a
	})
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func semicolon_merge(left, right []prix) []prix {
	ret := make([]prix, 0)
	lsize := len(left)
	rsize := len(right)
	var a, b int
	for i := 0; i < lsize; i++ {
		for j := 0; j < rsize; j++ {
			a = gcd(left[i].a, right[j].a)
			b = left[i].b + right[j].b
			if a != 0 {
				b = (b%a + a) % a
			}
			ret = append(ret, prix{a, b})
		}
	}
	return ret
}

func plus_merge(left, right []prix) []prix {
	sz := len(right)
	for i := 0; i < sz; i++ {
		left = append(left, right[i])
	}
	return left
}

func asterisk_merge(right []prix) []prix {
	ret := make([]prix, 0)
	if len(right) == 0 {
		return ret
	}
	g := gcd(right[0].a, right[0].b)
	sz := len(right)
	for i := 1; i < sz; i++ {
		g = gcd(g, gcd(right[i].a, right[i].b))
	}
	ret = append(ret, prix{g, 0})
	return ret
}

func minimize(a int, table [][]bool) {
	dvs := divisor_except_self(a)
	dvsz := len(dvs)
	for j := 0; j < dvsz; j++ {
		d := dvs[j]
		for k := 0; k < d; k++ {
			ok := true
			for l := 0; l < a/d; l++ {
				if !table[a][l*d+k] {
					ok = false
					break
				}
			}
			if ok {
				table[d][k] = true
			}
		}
	}
}

func divisor_except_self(p int) []int {
	ret := make([]int, 0)
	if p >= 1 {
		ret = append(ret, 1)
	}
	for i := 2; i*i <= p; i++ {
		if p%i == 0 {
			ret = append(ret, i)
			if i != p/i {
				ret = append(ret, i)
			}
		}
	}
	return ret
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type prix struct {
	a, b int
}

func reverseOrderPair(a []prix) []prix {
	n := len(a)
	res := make([]prix, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
