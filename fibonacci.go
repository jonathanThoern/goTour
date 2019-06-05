// Go functions may be closures. A closure is a function value that references
// variables from outside its body. The function may access and assign to the 
// referenced variables; in this sense the function is "bound" to the variables.

package main
import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	Fn := 0
	a := 0
	b := 1
	return func() int{

		//Fn, a, b = a, b, a+b
		// (Not really sure how this works)

        	Fn = a
        	a = a + b
		b = Fn
		return Fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}
