package main

import "fmt"

func main() {
	type pair struct {
		x, y int
	}

	const M = 2000

	var k int
	fmt.Scan(&k)
	res := make([]pair, 0)
	res = append(res, pair{M, 0})
	res = append(res, pair{M - 2, 0})
	res = append(res, pair{M - 2, 1})
	x, y := M-2, 1
	k -= 2
	for t := 0; 4*t+1 <= k; k, t = k-(4*t+1), t+1 {
		res = append(res, pair{-x, y})
		res = append(res, pair{-x, -y})
		x -= 2
		res = append(res, pair{x, -y})
		y += 2
		res = append(res, pair{x, y})
	}

	res = append(res, pair{1, y})
	y -= (k/2)*2 + 1
	res = append(res, pair{1, y})
	res = append(res, pair{0, y})
	res = append(res, pair{0, M})

	if (k & 1) != 0 {
		res = append(res, pair{M - 1, M})
		res = append(res, pair{M - 1, -M})
		res = append(res, pair{M, -M})
	} else {
		res = append(res, pair{M, M})
	}

	fmt.Println(len(res))
	for _, p := range res {
		fmt.Println(p.x, p.y)
	}
}
