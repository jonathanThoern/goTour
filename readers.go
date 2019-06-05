package main

import(
		"golang.org/x/tour/reader"
		"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error){

	for i := range b{
		b[i] = 65 // from ascii table (A)
	}

	// b is a reference and not a copy, will therefore be altered
	// (arrays are passed as pointers, same as in C?)
	// Return lenght and error status
	return len(b), nil
}

func main() {

	/** added for understanding **/
	b := make([]byte, 8)
	a := MyReader{}
	n, err := a.Read(b)
	for i, v := range b[:n]{
		fmt.Println("n:",n, "err:",err, "i:",i, "v:",v)
		fmt.Printf("char val: %c\n",v)
	} /************************************/

	reader.Validate(MyReader{})
}

