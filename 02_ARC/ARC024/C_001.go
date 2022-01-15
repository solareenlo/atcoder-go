package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	var s string
	fmt.Fscan(in, &n, &k, &s)

	cnt := [26]int{}
	for i := 0; i < k; i++ {
		cnt[s[i]-'a']++
	}

	m := map[[26]int]int{}
	m[cnt] = k - 1
	f := false
	for i := 0; i < n-k; i++ {
		cnt[s[i]-'a']--
		cnt[s[i+k]-'a']++
		if m[cnt] == 0 {
			m[cnt] = i + k
		} else if m[cnt] < i+1 {
			f = true
		}
	}

	if f {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
