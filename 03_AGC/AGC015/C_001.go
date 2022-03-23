package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	const MX = 2005
	s := [MX][MX]int{}
	for i := 1; i < n+1; i++ {
		var S string
		fmt.Fscan(in, &S)
		for j := range S {
			s[i][j+1] = int(S[j])
		}
	}

	ac := [MX][MX]int{}
	ac2 := [MX][MX]int{}
	ac3 := [MX][MX]int{}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			ac[i][j] = ac[i-1][j] + ac[i][j-1] - ac[i-1][j-1] + int(s[i][j]-'0')*int('1'-s[i-1][j]+'0'-s[i][j-1])
			ac2[i][j] = ac2[i][j-1] + int(s[i][j]-'0')*int(s[i-1][j]-'0')
			ac3[i][j] = ac3[i-1][j] + int(s[i][j]-'0')*int(s[i][j-1]-'0')
		}
	}

	for i := 1; i < q+1; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		fmt.Fprintln(out, ac[c][d]+ac[a-1][b-1]-ac[c][b-1]-ac[a-1][d]+ac2[a][d]-ac2[a][b-1]+ac3[c][b]-ac3[a-1][b])
	}
}
