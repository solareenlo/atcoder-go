package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	road := make([]int, 0)
	cnt := make([]int, n+1)
	for i := range cnt {
		cnt[i] = -1
	}

	tail, loop := 0, 1

	v := 1
	for cnt[v] == -1 {
		cnt[v] = len(road)
		road = append(road, v)
		v = a[v-1]
	}

	tail = cnt[v]
	loop = len(road) - tail

	if k < tail {
		fmt.Println(road[k])
	} else {
		k -= tail
		k %= loop
		fmt.Println(road[tail+k])
	}
}
