package basic

import (
	"errors"
)

func Handling(a, b int) (int, error){

	if b == 0 {
		return 0, errors.New("Error bro")
	}
	return a + b, nil
}