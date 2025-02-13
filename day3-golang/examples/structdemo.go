package main
import "fmt"
type student struct{
	Name string
	USN float64
	Dept string
}
func main(){
	st:=student{Name:"Meghana",USN:94.0,Dept:"CSE"}
	fmt.Println("Name: ",st.Name, "\nUSN: ",st.USN, "\nDept: ",st.Dept)
}