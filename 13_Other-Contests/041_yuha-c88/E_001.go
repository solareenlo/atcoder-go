package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	type pair struct {
		x, y int
	}

	var n int
	n = nextInt()

	s := make([][]string, 40000)
	var l, a [4]string
	var in [4]bool
	for i := 0; i < n; i++ {
		t := nextString()
		s[i] = utf8(t)
	}
	for i := 0; i < 5; i++ {
		t := nextString()
		tmp1 := utf8(t)
		if i == 0 {
			l[0] = tmp1[2]
		}
		if i == 1 {
			a[0] = tmp1[2]
		}
		if i == 2 {
			l[1] = tmp1[0]
			a[1] = tmp1[1]
			a[2] = tmp1[3]
			l[2] = tmp1[4]
		}
		if i == 3 {
			a[3] = tmp1[2]
		}
		if i == 4 {
			l[3] = tmp1[2]
		}
	}
	in[0] = a[0] == "↓"
	in[1] = a[1] == "→"
	in[2] = a[2] == "←"
	in[3] = a[3] == "↑"
	zip := make(map[string]int)
	unzip := make([]string, 0)
	for i := 0; i < n; i++ {
		unzip = append(unzip, s[i][0])
		unzip = append(unzip, s[i][1])
	}
	sort.Strings(unzip)
	unzip = unique(unzip)
	for i := 0; i < len(unzip); i++ {
		zip[unzip[i]] = i
	}
	sn := make(map[pair]bool)
	for i := 0; i < n; i++ {
		sn[pair{zip[s[i][0]], zip[s[i][1]]}] = true
	}

	var ln [4]int
	for i := 0; i < 4; i++ {
		ln[i] = zip[l[i]]
	}
	for i := 0; i < len(unzip); i++ {
		ok := true
		for j := 0; j < 4; j++ {
			var p pair
			if in[j] {
				p = pair{ln[j], i}
			} else {
				p = pair{i, ln[j]}
			}
			if _, OK := sn[p]; !OK {
				ok = false
			}
		}
		if ok {
			fmt.Println(unzip[i])
		}
	}
}

func utf8(s string) []string {
	ans := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if (s[i]>>7)&1 != 0 {
			ans = append(ans, s[i:i+3])
			i += 2
		} else {
			ans = append(ans, string(s[i]))
		}
	}
	return ans
}

func unique(a []string) []string {
	occurred := map[string]bool{}
	result := []string{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Strings(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

var scanner = bufio.NewScanner(os.Stdin)

func nextString() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	i, e := strconv.Atoi(nextString())
	if e != nil {
		panic(e)
	}
	return i
}
