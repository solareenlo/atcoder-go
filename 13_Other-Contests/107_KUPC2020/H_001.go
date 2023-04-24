package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1010

	var H, W int
	fmt.Fscan(in, &H, &W)
	var s [N]string
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &s[i])
	}
	var mex [40]int
	for i := 0; i < 33; i++ {
		for j := 0; j < 10; j++ {
			if (i & (1 << j)) == 0 {
				mex[i] = j
				break
			}
		}
	}
	var sg [N][N]int
	for i := 0; i < W; i++ {
		sg[H][i] = 4
	}
	var ls, rs, sl, sr, Go [N][5]int
	for i := H - 1; i >= 0; i-- {
		for j := 0; j < W; j++ {
			for fs := 0; fs < 5; fs++ {
				if s[i][j] == '#' {
					Go[j][fs] = 4
				} else {
					Go[j][fs] = mex[(1<<fs)|(1<<sg[i+1][j])]
				}
			}
		}
		for fs := 0; fs < 5; fs++ {
			s := fs
			for j := 0; j < W; j++ {
				ls[j][fs] = s
				s = Go[j][s]
			}
			s = fs
			for j := W - 1; j >= 0; j-- {
				rs[j][fs] = s
				s = Go[j][s]
			}
		}
		for fs := 0; fs < 5; fs++ {
			sl[0][fs] = Go[0][fs]
		}
		for j := 1; j < W; j++ {
			for fs := 0; fs < 5; fs++ {
				sl[j][fs] = sl[j-1][Go[j][fs]]
			}
		}
		for fs := 0; fs < 5; fs++ {
			sr[W-1][fs] = Go[W-1][fs]
		}
		for j := W - 1 - 1; j >= 0; j-- {
			for fs := 0; fs < 5; fs++ {
				sr[j][fs] = sr[j+1][Go[j][fs]]
			}
		}
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				sg[i][j] = 4
				continue
			}
			lsg := -1
			if W == 1 {
				lsg = 4
			} else if j != W-1 {
				lsg = ls[j][sr[j+1][4]]
			} else {
				lsg = ls[j][4]
			}
			rsg := -1
			if W == 1 {
				rsg = 4
			} else if j != 0 {
				rsg = rs[j][sl[j-1][4]]
			} else {
				rsg = rs[j][4]
			}
			sg[i][j] = mex[(1<<lsg)|(1<<rsg)|(1<<sg[i+1][j])]
		}
	}
	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == 'B' {
				ans ^= sg[i][j]
			}
		}
	}
	if ans != 0 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Bob")
	}
}
