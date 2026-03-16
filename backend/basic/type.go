package basic

import "fmt"

func TypeData() {

	name := "Bagas"
	age := 45
	grade := 3.9
	lecture := [3]int{1,2,3} // Array
	slice := []int{1,2,3} // Slice

	fmt.Printf("%T\n",name)
	fmt.Printf("%T\n",age)
	fmt.Printf("%T\n",grade)
	fmt.Printf("%T\n",lecture)
	fmt.Printf("%T\n",slice)

}