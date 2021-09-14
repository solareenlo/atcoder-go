package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Scan(&s[0], &s[1], &s[2])

	i := 0
	for len(s[i]) > 0 {
		j := s[i][0] - 'a'
		s[i] = s[i][1:]
		i = int(j)
	}
	fmt.Println(string('A' + byte(i)))
}
