/* Regarding types */
import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Will print the type
func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
/********************************************************/


/* Bitwize creating big numbers (numeric constants - high precision values)*/
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
/*******************************************************/

/* Go has ONLY for-loops (this is equivalent to a while(?), init and post statements are optional)*/
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
		fmt.Println(sum)
	}
}

//The "While-loop in Go convention (semicolons can be left out)
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// For endless while:
func main() {
	for {
		//Do something endlessly
	}
}

/*****************************************************/

/* Deferred functions will execute after others. Pushed to stack and executed LIFO */

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // prints last in reverse order (since LIFO)
	}

	fmt.Println("done")
}
/********************************************************/

/* Arrays and SLICES*/

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Slices are more common than arrays (dynamic size)
//A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	a[low : high]
//This selects a half-open range which includes the first element, but excludes the last one.
// May also be omitted as:
	a[0:10]
	a[:10] //from first to 9th
	a[0:] //from 0 to lenght of
	a[:] //all

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	s = primes[1:6]
	fmt.Println(s)
}

//A slice does not store any data, it just describes a section of an underlying array.
//Changing the elements of a slice modifies the corresponding elements of its underlying array.
//Other slices that share the same underlying array will see those changes.
// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// declaring and initializing a struct array
s := []struct {
	i int
	b bool
}{
	{2, true},
	{3, false},
	{5, true},
	{7, true},
	{11, false},
	{13, true},
}

// Very hard to understand...
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//The make function allocates a zeroed array and returns a slice that refers to that array:
a := make([]int, 5)  // len(a)=5
//To specify a capacity, pass a third argument to make:
b := make([]int, 0, 5) // len(b)=0, cap(b)=5


// Append to extend slices (will re-allocate memory if needed)
func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// Range over slice
// two values are returned for each iteration, index and a copy of the element at that index.

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// skip the index or value by assigning to _.

for i, _ := range pow
for _, value := range pow

//If only want index, omit the second variable.

for i := range pow

// As in example:

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// Slices of slices example
func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/********************************************************/


/* MAPS */


// Basic map
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Test Location"] = Vertex{
		68.12345, -86.54321,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Test Location"])
	fmt.Println(m)
}
// fmt.Println(m) gives:
// map[Bell Labs:{40.68433 -74.39967} Test Location:{68.12345 -86.54321}]


// Map Literals are like struct literals - Keys required
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

// If top-level is just type name it can be omitted from elements of literal
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
// example
fmt.Println(m["Google"].Long) // prints -122.08408

// Mutating Maps:

//Insert or update an element in map m:
m[key] = elem

//Retrieve an element:
elem = m[key]

//Delete an element:
delete(m, key)

//Test that a key is present with a two-value assignment:
elem, ok = m[key]

//If key is in m, ok is true. If not, ok is false.
//If key is not in the map, then elem is the zero value for the map's element type.

//Note: If elem or ok have not yet been declared you could use a short declaration form:
elem, ok := m[key]


/* Function Values */

//Functions are values too. They can be passed around just like other values.
//Function values may be used as function arguments and return values.


// Try to understand this...
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

//Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

//For example, the adder function returns a closure. Each closure is bound to its own sum variable

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(5),
			neg(-2*i),
		)
	}
}

/****************************************************/

/* METHODS */

// A method is a function with a special receiver argument.
//The receiver appears in its own argument list between the func keyword and the method name.
//In this example, the Abs method has a receiver of type Vertex named v.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Sum() float64 {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.Sum())
}


// You can declare a method on non-struct types, too.
//In this example we see a numeric type MyFloat with an Abs method.

//You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

type MyFloat float64
type MyString string

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (s MyString) Lenght() int{
	return len([]rune(s))
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	s := MyString("You go now!")
	fmt.Println(s.Lenght())
}


// Pointer receivers

//Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { // Changes value since pointer
	v.X = v.X * f
	v.Y = v.Y * f
}

// while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
// For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

// while methods with value receivers take either a value or a pointer as the receiver when they are called:
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
// In this case, the method call p.Abs() is interpreted as (*p).Abs().


/* INTERFACES */

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Interfaces are implemented implicitly
// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

//Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

// Interface values
// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


//If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
//In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
//Note that an interface value that holds a nil concrete value is itself non-nil.

func main() {
	var i I

	var t *T
	i = t
	describe(i) // gives: (<nil>, *main.T)
	i.M()

	i = &T{"hello", 10}
	describe(i) // gives: (&{hello 10}, *main.T)
	i.M()
}

// A nil interface value holds neither value nor concrete type.
// Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

func main() {
	var i I
	describe(i) // gives: (<nil>, <nil>)
	i.M() // gives runtime error
}

// The interface type that specifies zero methods is known as the empty interface:
interface{}
// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

func main() {
	var i interface{}
	describe(i)	// gives: (<nil>, <nil>)

	i = 42
	describe(i)	// gives: (42, int)

	i = "hello"
	describe(i)	// gives: (hello, string)
}


/* Type assertions */

// A type assertion provides access to an interface value's underlying concrete value.
t := i.(T)
// To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
t, ok := i.(T)
// If i holds a T, then t will be the underlying value and ok will be true. 
// (similar reading from a map)

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)		// gives: hello

	s, ok := i.(string)
	fmt.Println(s, ok)	// gives: hello true	

	f, ok := i.(float64)
	fmt.Println(f, ok) 	// gives: 0 false

	f = i.(float64) // panic (i.e program crashes)
	fmt.Println(f)
}

// Type switches

//A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
//The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.
//  In the default case (where there is no match), the variable v is of the same interface type and value as i.

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)		// gives first case
	do("hello")	// gives second case
	do(true)	// gives third case
}

// Stringers

//One of the most ubiquitous interfaces is Stringer defined by the fmt package.
type Stringer interface {
    String() string
}
//A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.


/* ERRORS */

// Go programs express error state with error values.

// The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}
// (As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
// A nil error denotes success; a non-nil error denotes failure.

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil { // unclear how this works with dual conditions
		fmt.Println(err) // gives: at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
	}
}

/**************************************************************************************/


/* READERS */

// The io package specifies the io.Reader interface, which represents the read end of a stream of data.
//The Go standard library contains many implementations of these interfaces, including files, network connections, compressors, ciphers, and others.
//The io.Reader interface has a Read method:

func (T) Read(b []byte) (n int, err error)

//Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

//The example code creates a strings.Reader and consumes its output 8 bytes at a time.

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}


/* IMAGES */

// Package image defines the Image interface:

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}

// Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

// The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package


/*************************************************************/

/* GOROUTINES */


// A goroutine is a lightweight thread managed by the Go runtime.

go f(x, y, z)
// starts a new goroutine running

f(x, y, z)
// The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

// Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. 

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}


/* Channels */

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
// (The data flows in the direction of the arrow.)

// Like maps and slices, channels must be created before use:
ch := make(chan int)

// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

// The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

//Buffered Channels

//Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
ch := make(chan int, 100)

//Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

//Modify the example to overfill the buffer and see what happens.

func main() {
	ch := make(chan int, 2)
	
	ch <- 1
	ch <- 2
	ch <- 3 	// adding third causes deadlock
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // causes deadlock
}

// Range and Close

// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.

// The loop 
for i := range c 
// receives values from the channel repeatedly until it is closed.

// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
//Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// Select

//The select statement lets a goroutine wait on multiple communication operations.
//A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// Default Selection

// The default case in a select is run if no other case is ready.
// Use a default case to try a send or receive without blocking:
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}

// example:

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/* Binary trees (exercise) */

// A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

// This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}


/** sync.MUTEX **/


// Go's standard library provides mutual exclusion with sync.Mutex and its two methods:
// Lock
// Unlock
//We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.
// We can also use defer to ensure the mutex will be unlocked as in the Value method.


import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
























