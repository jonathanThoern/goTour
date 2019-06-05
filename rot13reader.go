// A common pattern is an io.Reader that wraps another io.Reader, modifying
// the stream in some way.
// For example, the gzip.NewReader function takes an io.Reader (a stream of
// compressed data) and returns a *gzip.Reader that also implements io.Reader
// (a stream of the decompressed data).
// Implement a rot13Reader that implements io.Reader and reads from an
// io.Reader, modifying the stream by applying the rot13 substitution cipher
// to all alphabetical characters.

package main

import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

type ErrRead int

func (e ErrRead) Error() string{
	return fmt.Sprintf("Something went wrong")
}

func (rot rot13Reader) Read(b []byte) (int, error){
	n, err := rot.r.Read(b)
	if err != nil{
		return 0, ErrRead(n)
	}
	//fmt.Println(n,err)
	for i := range b[:n]{
		if (b[i] >= 65 && b[i] <= 77) || (b[i] >= 97 && b[i] <= 109){
			b[i] += 13
		} else if (b[i] >= 78 && b[i] <= 90) || (b[i] >= 110 && b[i] <= 122){
			b[i] -= 13
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
