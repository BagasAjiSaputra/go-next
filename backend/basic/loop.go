package basic

import "fmt"

func Loopink() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}

func Whiles() {

	i := 0

	for i < 5 {
		i++
		fmt.Println("Loop", i)
	}

}

func Ranges() {

	listink := []int{1,2,3,4,5,6}

	for i, value := range listink {
		fmt.Println("Index :", i,"Loop :", value)
	}

}