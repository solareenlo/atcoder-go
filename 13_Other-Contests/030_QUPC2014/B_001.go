package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	sd := [15]string{"Nil", "Un", "Bi", "Tri", "Quad", "Pent", "Hex", "Sept", "Oct", "Enn"}
	sx := [15]string{"nil", "un", "bi", "tri", "quad", "pent", "hex", "sept", "oct", "enn"}
	fmt.Printf("%s", sd[s[0]-'0'])
	if s[1] == '0' && s[0] != '9' {
		fmt.Print("nil")
	} else if s[1] == '0' && s[0] == '9' {
		fmt.Print("il")
	} else {
		fmt.Printf("%s", sx[s[1]-'0'])
	}

	if s[2] == '0' && s[1] != '9' {
		fmt.Print("nil")
	} else if s[2] == '0' && s[1] == '9' {
		fmt.Print("il")
	} else {
		fmt.Printf("%s", sx[s[2]-'0'])
	}

	if s[2] == '2' || s[2] == '3' {
		fmt.Print("um")
	} else {
		fmt.Print("ium")
	}
	fmt.Println()
}
