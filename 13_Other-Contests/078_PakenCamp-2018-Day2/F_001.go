package main

import "fmt"

func main() {
	type pair struct {
		x, y int
	}

	var k int
	fmt.Scan(&k)
	v := make([]pair, 0)
	for i := 2; i <= 5; i++ {
		v = append(v, pair{1, i})
	}
	for i := 2; i <= 114; i += 4 {
		for j := 4; j <= 7; j++ {
			if i == 114 && j == 7 {
				continue
			}
			v = append(v, pair{i, i + j})
			v = append(v, pair{i + 1, i + j})
			v = append(v, pair{i + 2, i + j})
			v = append(v, pair{i + 3, i + j})
		}
	}
	for i := 121; i <= 149; i++ {
		v = append(v, pair{i, i + 1})
	}
	y := 121
	j := 2
	for k > 0 {
		x := k % 4
		k /= 4
		for i := 0; i < x; i++ {
			v = append(v, pair{j + i, y})
		}
		j += 4
		y++
	}
	fmt.Println(150, len(v))
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i].x, v[i].y)
	}
}
