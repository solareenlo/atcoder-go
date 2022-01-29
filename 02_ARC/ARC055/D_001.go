package main

import (
	"fmt"
	"math"
	"sort"
)

const eps = 1e-9

func GO(x, y int) int {
	return (y - x + 10) % 10
}

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	m := map[int]bool{}
	for i := 0; i < n-1; i++ {
		m[GO(int(s[i]), int(s[i+1]))] = true
	}
	if len(m) < 2 {
		fmt.Println(string(s[0]))
		return
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	if len(m) > 2 || (GO(keys[0], keys[len(keys)-1]) != 1 && GO(keys[len(keys)-1], keys[0]) != 1) {
		fmt.Println(-1)
		return
	}
	var d int
	if keys[len(keys)-1] == keys[0]+1 {
		d = keys[0]
	} else {
		d = keys[len(keys)-1]
	}
	a := [10010]int{}
	a[0] = int(s[0] - '0')
	for i := 1; i < n; i++ {
		a[i] = a[i-1] + d
		if a[i]%10 != int(s[i]-'0') {
			a[i]++
		}
		if a[i]%10 != int(s[i]-'0') {
			fmt.Println(-1)
			return
		}
	}
	maxi := 0.0
	mini := 1e9
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			maxi = math.Max(maxi, (float64(a[j]-(a[i]+1))+eps)/float64(j-i))
			mini = math.Min(mini, (float64((a[j]+1)-a[i])-eps)/float64(j-i))
		}
	}
	if maxi > mini {
		fmt.Println(-1)
		return
	}
	pw := 1
	for math.Ceil(maxi*float64(pw)) > math.Floor(mini*float64(pw)) {
		pw *= 10
	}
	B := math.Floor(mini * float64(pw))
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, a[i]*pw-int(B)*i)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
