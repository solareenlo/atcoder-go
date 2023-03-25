package main

import "fmt"

func main() {
	fmt.Println("a[0]=1201")
	fmt.Println("a[9990]=1")
	fmt.Println("while(a[9990]<401){")
	fmt.Println("a[a[9990]+a[9990]]=2")
	fmt.Println("a[a[9990]+800]=a[9990]+a[9990]-1")
	fmt.Println("a[9990]=a[9990]+1")
	fmt.Println("}")
	fmt.Println("----")
	fmt.Println("while(0<a[a[X+800]]){")
	fmt.Println("a[9990]=a[a[X+800]]")
	fmt.Println("a[a[X+800]]=a[0]")
	fmt.Println("a[a[0]+a[9990]]=a[9990]+a[9990]")
	fmt.Println("a[X+800]=a[0]")
	fmt.Println("a[0]=a[0]+a[9990]+1")
	fmt.Println("")
	fmt.Println("}")
	fmt.Println("a[a[X+800]]=Y")
	fmt.Println("a[X+800]=a[X+800]+1")
	fmt.Println("----")
	fmt.Println("a[9990]=X+X-1")
	fmt.Println("a[9991]=1")
	fmt.Println("a[9992]=Y")
	fmt.Println("while(a[9991]<a[9992]){")
	fmt.Println("a[9990]=a[a[9990]+a[9991]]")
	fmt.Println("a[9992]=a[9992]-a[9991]")
	fmt.Println("a[9991]=a[9991]+a[9991]")
	fmt.Println("}")
	fmt.Println("A=a[a[9990]+a[9992]-1]")
}
