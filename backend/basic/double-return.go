package basic

import "fmt"

func Divides(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cant Zero Bruh")
	}

	return a / b, nil
}