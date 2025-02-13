package main

import "fmt"
func main() {
    ifelseDemo()
}
func ifelseDemo(){
	var num1,num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	if num1 > num2{
		fmt.Println("num1 is larger")
	}
	else if num1 ==num2{
		fmt.Println("num1 and num2 are equal")
	}else{
		fmt.Println("num2 is larger")
	}
}


