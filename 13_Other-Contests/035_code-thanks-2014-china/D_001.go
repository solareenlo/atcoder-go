package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var H, W int
var S [51][]string
var D [3][51][51]int
var X, Y [3]int
var memo [51][51][51][51]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &H, &W)

	for i := range D {
		for j := range D[i] {
			for k := range D[i][j] {
				D[i][j][k] = 1061109567
			}
		}
	}
	Q := make([]int, 0)
	for i := range S {
		S[i] = make([]string, 59)
	}
	for y := 0; y < H; y++ {
		var s string
		fmt.Fscan(in, &s)
		s += strings.Repeat("", 59)
		S[y] = strings.Split(s, "")
		for x := 0; x < W; x++ {
			if S[y][x] == "S" {
				X[0] = x
				Y[0] = y
				D[0][y][x] = 0
				Q = append(Q, 0+y*100+x)
			}
			if S[y][x] == "A" {
				X[1] = x
				Y[1] = y
				D[1][y][x] = 0
				Q = append(Q, 10000+y*100+x)
			}
			if S[y][x] == "B" {
				X[2] = x
				Y[2] = y
				D[2][y][x] = 0
				Q = append(Q, 20000+y*100+x)
			}
		}
	}

	for len(Q) > 0 {
		k := Q[0]
		Q = Q[1:]
		st := k / 10000
		cy := k / 100 % 100
		cx := k % 100
		for i := 0; i < 4; i++ {
			dd := [4]int{0, 1, 0, -1}
			tx := cx + dd[i]
			ty := cy + dd[i^1]
			if tx < 0 || ty < 0 || tx >= W || ty >= H {
				continue
			}
			if S[ty][tx] == "#" {
				continue
			}
			if D[st][ty][tx] <= D[st][cy][cx]+1 {
				continue
			}
			D[st][ty][tx] = D[st][cy][cx] + 1
			Q = append(Q, st*10000+ty*100+tx)
		}
	}

	for y := 0; y < 51; y++ {
		for x := 0; x < 51; x++ {
			for i := 0; i < 51; i++ {
				for j := 0; j < 51; j++ {
					memo[y][x][i][j] = -2
				}
			}
		}
	}

	if ok(X[0], Y[0], X[0], Y[0]) < 0 {
		fmt.Println("NA")
	} else {
		fix(X[0], Y[0], X[0], Y[0])
		for y := 0; y < H; y++ {
			fmt.Println(strings.Join(S[y], ""))
		}
	}
}

func ok(x1, y1, x2, y2 int) int {
	if D[1][y1][x1] == 0 && D[2][y2][x2] == 0 {
		return 0
	}

	if memo[x1][y1][x2][y2] != -2 {
		return memo[x1][y1][x2][y2]
	}
	memo[x1][y1][x2][y2] = -1

	var turn int
	if D[0][y1][x1] != D[0][y2][x2] {
		turn = 1
	} else {
		turn = 0
	}
	if D[1][y1][x1] == 0 {
		turn = 1
	}
	if D[2][y2][x2] == 0 {
		turn = 0
	}

	for i := 0; i < 4; i++ {
		dd := [4]int{0, 1, 0, -1}
		var tx, ty int
		if turn == 0 {
			tx = x1 + dd[i]
			ty = y1 + dd[i^1]
			if tx < 0 || ty < 0 || tx >= W || ty >= H {
				continue
			}
			if tx == x2 && ty == y2 {
				continue
			}
			if D[1][ty][tx] != D[1][y1][x1]-1 {
				continue
			}
			if ok(tx, ty, x2, y2) >= 0 {
				memo[x1][y1][x2][y2] = i
				return memo[x1][y1][x2][y2]
			}
		} else {
			tx = x2 + dd[i]
			ty = y2 + dd[i^1]
			if tx < 0 || ty < 0 || tx >= W || ty >= H {
				continue
			}
			if tx == x1 && ty == y1 {
				continue
			}
			if D[2][ty][tx] != D[2][y2][x2]-1 {
				continue
			}
			if ok(x1, y1, tx, ty) >= 0 {
				memo[x1][y1][x2][y2] = i
				return memo[x1][y1][x2][y2]
			}
		}
	}

	return memo[x1][y1][x2][y2]
}

func fix(x1, y1, x2, y2 int) {
	turn := 0

	if D[1][y1][x1] == 0 && D[2][y2][x2] == 0 {
		return
	}

	if D[0][y1][x1] != D[0][y2][x2] {
		turn = 1
	} else {
		turn = 0
	}
	if D[1][y1][x1] == 0 {
		turn = 1
	}
	if D[2][y2][x2] == 0 {
		turn = 0
	}

	dd := [4]int{0, 1, 0, -1}
	if turn == 0 && S[y1][x1] == "." {
		S[y1][x1] = "a"
	}
	if turn == 1 && S[y2][x2] == "." {
		S[y2][x2] = "b"
	}
	if turn == 0 {
		fix(x1+dd[memo[x1][y1][x2][y2]], y1+dd[memo[x1][y1][x2][y2]^1], x2, y2)
	}
	if turn == 1 {
		fix(x1, y1, x2+dd[memo[x1][y1][x2][y2]], y2+dd[memo[x1][y1][x2][y2]^1])
	}
}
