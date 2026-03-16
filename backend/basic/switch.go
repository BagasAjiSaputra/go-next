package basic

import "fmt"

func Switches() {

	grade := "C"

	switch grade {
	case "A":
		fmt.Println("Good")
	case "B":
		fmt.Println("Enough")
	case "C":
		fmt.Println("Try Tomorrow")
	}

}