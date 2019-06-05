package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// From earlier exercise (changed return)
func Sqrt(x float64) (float64, error) {

	//error handling
	if(x < 0){
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	prev := float64(0)
	for i := 1 ; ; i++{
		z -= (z*z - x) / (2*z)
		//fmt.Println(i, z)
		if math.Abs(prev - z) < 0.000000000000001{
			return z, nil
		}
		prev = z
	}
}

func main() {
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(-2))
}
