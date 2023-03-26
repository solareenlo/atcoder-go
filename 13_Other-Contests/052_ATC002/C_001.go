package main

import "fmt"

var arr [200003]int
var w [200003]int
var st int
var ans int64

func insert(c, e int) {
	var i, d int
	ans = ans + int64(c)
	st--
	for i = e; i <= st; i++ {
		arr[i] = arr[i+1]
	}
	for i = e - 2; i >= 0; i-- {
		if c <= arr[i] {
			arr[i+1] = c
			break
		}
		arr[i+1] = arr[i]
	}
	for i > 1 && arr[i-1] <= c {
		d = st - i
		insert(arr[i-1]+arr[i], i)
		i = st - d
	}
}

func main() {
	var i, n int
	ans = 0
	st = 0
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&w[i])
	}
	w[0] = 2000000002
	w[n+1] = 2000000001
	st++
	arr[st] = w[0]
	st++
	arr[st] = w[1]
	for i = 2; i <= n+1; i++ {
		for arr[st-1] <= w[i] {
			insert(arr[st-1]+arr[st], st)
		}
		st++
		arr[st] = w[i]
	}
	fmt.Println(ans)
}
