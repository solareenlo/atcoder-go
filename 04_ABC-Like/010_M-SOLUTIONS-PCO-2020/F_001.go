package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	N int
	X = [1 << 18]int{}
	Y = [1 << 18]int{}
	D = [1 << 18]string{}
)

func Rotate() {
	for i := 1; i <= N; i++ {
		cx := Y[i]
		cy := 200000 - X[i]
		var cz string
		if D[i] == "U" {
			cz = "R"
		}
		if D[i] == "R" {
			cz = "D"
		}
		if D[i] == "D" {
			cz = "L"
		}
		if D[i] == "L" {
			cz = "U"
		}
		X[i] = cx
		Y[i] = cy
		D[i] = cz
	}
}

type pair struct {
	x int
	y string
}

var (
	V1 = [1 << 19][]pair{}
	V2 = [1 << 19][]pair{}
)

func solve_01() int {
	Answer := 1 << 24
	for i := 0; i <= 500000; i++ {
		V1[i] = V1[i][:0]
	}
	for i := 1; i <= N; i++ {
		if !(D[i] == "R" || D[i] == "U") {
			continue
		}
		V1[X[i]+Y[i]] = append(V1[X[i]+Y[i]], pair{X[i], D[i]})
	}
	for i := 0; i <= 500000; i++ {
		sort.Slice(V1[i], func(a, b int) bool {
			return V1[i][a].x < V1[i][b].x
		})
		for j := 0; j < len(V1[i])-1; j++ {
			if !(V1[i][j].y == "R" && V1[i][j+1].y == "U") {
				continue
			}
			Answer = min(Answer, V1[i][j+1].x-V1[i][j].x)
		}
	}
	return Answer * 10
}

func solve_02() int {
	Answer := 1 << 24
	for i := 0; i <= 500000; i++ {
		V2[i] = V2[i][:0]
	}
	for i := 1; i <= N; i++ {
		if !(D[i] == "L" || D[i] == "R") {
			continue
		}
		V2[Y[i]] = append(V2[Y[i]], pair{X[i], D[i]})
	}
	for i := 0; i <= 500000; i++ {
		sort.Slice(V2[i], func(a, b int) bool {
			return V2[i][a].x < V2[i][b].x
		})
		for j := 0; j < len(V2[i])-1; j++ {
			if !(V2[i][j].y == "R" && V2[i][j+1].y == "L") {
				continue
			}
			Answer = min(Answer, V2[i][j+1].x-V2[i][j].x)
		}
	}
	return Answer * 5
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &X[i], &Y[i], &D[i])
	}

	FinalAns := 1 << 24
	for t := 1; t <= 4; t++ {
		val1 := solve_01()
		FinalAns = min(FinalAns, val1)
		val2 := solve_02()
		FinalAns = min(FinalAns, val2)
		Rotate()
	}
	if FinalAns == 1<<24 {
		fmt.Println("SAFE")
	} else {
		fmt.Println(FinalAns)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
