package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b int
	fmt.Fscan(in, &a, &b)
	mp := make(map[int]string)
	s := strings.Split("???????", "")
	var t string
	for {
		hash := 0
		for i := 0; i < 7; i++ {
			s[i] = string(byte(xor128()%26) + 'a')
			hash = (hash*a + int(s[i][0]-'a') + 1) % b
		}
		if _, ok := mp[hash]; !ok {
			mp[hash] = strings.Join(s, "")
		} else if strings.Join(s, "") != mp[hash] {
			t = mp[hash]
			break
		}
	}
	S := strings.Join(s, "")
	for i := 0; i < 100; i++ {
		var res string
		for j := 0; j < 7; j++ {
			if ((i >> j) & 1) != 0 {
				res += S
			} else {
				res += t
			}
		}
		fmt.Fprintln(out, res)
	}
}

var x = 123456789
var y = 362436069
var z = 521288629
var w = 88675123

func xor128() int {
	t := x ^ (x << 11)
	x = y
	y = z
	z = w
	w = (w ^ (w >> 19)) ^ (t ^ (t >> 8))
	return w
}
