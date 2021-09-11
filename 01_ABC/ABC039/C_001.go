package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	s += s

	cnt := 0
	for i := 0; i < len(s)-5; i++ {
		if s[i+1] == 'B' && s[i+3] == 'B' && s[i+5] == 'B' {
			cnt = i
			break
		}
	}

	scale := make([]string, 12)
	scale[5] = "Do"
	scale[3] = "Re"
	scale[1] = "Mi"
	scale[0] = "Fa"
	scale[10] = "So"
	scale[8] = "La"
	scale[6] = "Si"
	fmt.Println(scale[cnt%12])
}
