package main

import "fmt"

var s string
var lenmemo, idmemo [10101]int

func main() {
	var B, L, N int
	fmt.Scan(&B, &L, &N, &s)
	id := 0
	M := Len(&id)
	if B < 0 {
		B += M
	}
	id = 0
	fmt.Println(cut(id, B, L))
}

func Len(id *int) int {
	t := id
	if lenmemo[*id] != 0 {
		*id = idmemo[*t]
		return lenmemo[*t]
	}
	ret := 0
	for *id < len(s) && s[*id] != ')' {
		if s[*id] == '(' {
			*id++
			t := Len(id)
			*id++
			k := 0
			for s[*id] >= '0' && s[*id] <= '9' {
				k = k*10 + int(s[*id]-'0')
				*id++
				if *id >= len(s) {
					break
				}
			}
			ret += t * k
		} else {
			ret++
			*id++
		}
	}
	lenmemo[*t] = ret
	idmemo[*t] = *id
	return ret
}

func cut(id, B, L int) string {
	ret := ""
	for L > 0 && id < len(s) && s[id] != ')' {
		if s[id] == '(' {
			id++
			jd := id
			l := Len(&jd)
			jd++
			k := 0
			for s[jd] >= '0' && s[jd] <= '9' {
				k = k*10 + int(s[jd]-'0')
				jd++
				if jd >= len(s) {
					break
				}
			}
			if k*l <= B {
				B -= k * l
			} else if L <= l {
				t := B / l
				k -= t + 1
				B -= t * l
				x := cut(id, B, L)
				B = 0
				L -= len(x)
				ret += x
				if k > 0 {
					ret += cut(id, 0, L)
					L = 0
				}
			} else {
				t := B / l
				k -= t + 1
				B -= t * l
				x := cut(id, 0, l)
				tmp := l - B
				if L < l-B {
					tmp = L
				}
				xx := x[B : B+tmp]
				B = 0
				L -= len(xx)
				ret += xx
				for L >= l && k > 0 {
					ret += x
					L -= l
					k--
				}
				if k > 0 {
					ret += x[:L]
					L = 0
				}
			}
			id = jd
		} else {
			if B > 0 {
				B--
			} else {
				L--
				ret += string(s[id])
			}
			id++
		}
	}
	return ret
}
