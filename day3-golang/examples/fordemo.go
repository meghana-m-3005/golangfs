package main

import "fmt"
func main() {
    forPythonStyle()
}
func forCountDemo(){
	n := 1
	for n < 5{
		n *= 2
	}
	fmt.Println(n)
}

func forPythonStyle(){
	strings := []string{"hello","world","golang","NIE"}
	for _, s:=range strings{
		fmt.Println(s)
	}
}


