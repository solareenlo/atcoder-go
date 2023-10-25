package main

import (
	"fmt"
	"sort"
)

func main() {
	var pr [111]int

	var a int
	fmt.Scan(&a)
	if a == 1 {
		fmt.Println(1)
		fmt.Println("+", 1)
		return
	}
	now := 0.0
	pr[0] = -1
	pr[1] = -1
	v := make([]pair, 0)
	p := make([]int, 0)
	for i := 2; i <= a; i++ {
		if pr[i] == 0 {
			pr[i] = 1
			for j := i + i; j <= a; j += i {
				pr[j] = -1
			}
			F := i
			for F*i <= a {
				F *= i
			}
			p = append(p, F)
		}
	}
	for i := 0; i < len(p); i++ {
		tmp := 1
		for j := 0; j < len(p); j++ {
			if i == j {
				continue
			}
			tmp = tmp * p[j] % p[i]
		}
		req := 0
		val := 0
		for val != 1 {
			req++
			val = (val + tmp) % p[i]
		}
		if req*2 < p[i] {
			now += float64(req) / float64(p[i])
			for j := 0; j < req; j++ {
				v = append(v, pair{1, p[i]})
			}
		} else {
			now -= float64(p[i]-req) / float64(p[i])
			for j := 0; j < p[i]-req; j++ {
				v = append(v, pair{-1, p[i]})
			}
		}
	}
	for now < -0.2 {
		now += 1
		v = append(v, pair{1, 1})
	}
	for now > 0.8 {
		now -= 1
		v = append(v, pair{-1, 1})
	}
	sortPair(v)
	reverseOrderPair(v)
	fmt.Println(len(v))
	for i := 0; i < len(v); i++ {
		if v[i].x == 1 {
			fmt.Print("+ ")
		} else {
			fmt.Print("- ")
		}
		fmt.Println(v[i].y)
	}
}

func reverseOrderPair(a []pair) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
