package tasks

// this style of importing is also called factored import statement
import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"

	"golang.org/x/tour/wc"
)

// Some Global Variables without initialisers
var c, python, java bool

// Global variables with initialisation
var i, j int = 1, 2

// Constants can be character, string, boolean, or numeric values.
// Constant cannot be declared using := syntax
const (
	Pi    = 3.14
	Truth = true
)

// Person struct to define the person
type Person struct {
	Name   string
	Age    int
	isMale bool
}

// Hello function is just a starting function
func Hello() {
	fmt.Println("Hello World!")
}

// Time func prints current time
func Time() {
	fmt.Println("The Time is", time.Now())
}

// Maths does some math operations
func Maths() {
	// To seed the deterministic state
	rand.Seed(int64(time.Now().Second()))
	fmt.Println("Fav No is ", rand.Intn(10))

	fmt.Printf("Now you have %g problems \n", math.Sqrt(7))

	fmt.Println("I am Pi", math.Pi)

}

func checkEvenOdd(x int) string {
	if x%2 == 0 {
		fmt.Printf("%v is Even \n", x)
		return "Even"
	}
	fmt.Printf("%v is Odd \n", x)
	return "Odd"
}

// SomeFunc is a function syntax
// You can combine the arguments if they have same types
// You can also give named return values to make the code more readable.
func SomeFunc(x int, y int) (a, b string) {
	a = checkEvenOdd(x)
	b = checkEvenOdd(y)
	return
}

// VarDeclaration shows ways to declare variables using short syntax also shows Type Inference
func VarDeclaration() {
	someInt := 3
	someString := "Hi"
	someBool := true

	fmt.Println(someBool, someInt, someString)
}

// Basic Go Types

// bool defaults to false

// string defaults to ""

// int defaults to 0

// byte alias for uint8

// rune alias for int32, also represents unicode code point

// float32 float64

// complex64 complex128

// TypeConversion shows how to do explicit typecasting.
func TypeConversion() {
	a := 10
	b := 15.67
	c := float64(a) + b

	fmt.Println(c)
}

// ForLoop syntax
func ForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//ForLoopRevisited shows init and post statements are optional
func ForLoopRevisited() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// ForInfinite do not run
func ForInfinite() {
	for {

	}
}

// ShortIf x to power n to limit lim
func ShortIf(x, n, lim float64) float64 {

	// here v is only available to if and its else statements.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// Cant use v here now.
	return lim
}

// Sqrt is Custom Square Root Function
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return z
}

// LearnSwitch to learn the statement syntax
func LearnSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

// LongIfElse using switch without condition
func LongIfElse() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening!")
	}
}

// DeferFunc to see how defer works
func DeferFunc() {
	defer fmt.Print("World")
	fmt.Print("Hello")
}

// DeferStack works in LIFO basis
func DeferStack() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

// Pointers Explanation
func Pointers() {
	i, j := 42, 2701

	p := &i // point to i
	// fmt.Printf("%d", p)
	fmt.Println(*p) // read i through the pointer (dereferencing)
	fmt.Println(p)  // print reference
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// PersonStruct to define a person
func PersonStruct() {
	p := Person{"Harsh", 21, true} // Name:"Harsh", Age:21, IsMale: true
	q := &p

	q.Name = "John" // Struct fields can be accessed through struct pointer
	fmt.Println(p)
	fmt.Println(q)

	p1 := Person{} // Name:"", Age:0, IsMale: false
	fmt.Println(p1)

	p2 := Person{Name: "Harsh"} // Age and IsMale implicit to default values
	fmt.Println(p2)

}

// SomeArray syntax
func SomeArray() {
	var anArray [2]string // [n]T array of n values of type T
	// Array cannot be resized
	anArray[0] = "hello"
	anArray[1] = "World"

	fmt.Println(anArray)
}

//SliceExample use
func SliceExample() {

	t := make([]int, 5, 20) // make slice of int with length 5 capacity 20
	fmt.Printf("%d %d %d", t, len(t), cap(t))
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	// Zero value of slice is nil hich has a capacity and length of 0
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil")
	}

	var a []int

	// append on nil slice
	a = append(a, 0)

	// Slice grows as needed.
	a = append(a, 1)

	// add multiple elements at a time.
	a = append(a, 2, 3, 4)
}

// RangeExample usecase
func RangeExample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// when ranging over the slice, two values are returned, index and the COPY of the element at that index.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// drop , value if value is not needed
	for i := range pow {
		fmt.Printf("%d\n", i)
	}

	// if you dont want index put _ inplace of i
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

// ShowPic to show the pic
// func ShowPic() {
// 	pic.Show(Pic)
// }

// MapExample usecase
func MapExample() {
	// declare and add separately
	var p = make(map[string]Person)
	p["First"] = Person{
		"Harsh",
		21,
		true,
	}

	// declare and initialize
	var q = map[string]Person{
		"First":  {"Harsh", 21, true},
		"Second": {"Abhi", 24, true},
	}

	p["First"] = Person{"XYZ", 32, false} // change the value

	fmt.Println(p["First"], q)

	delete(p, "First") // delete the element

	v, ok := p["First"] //check if "First" key is present in p
	// Yes then ok will be true otherwise false.
	// If key is not in p then the value of the v will be default value of its type
	// which would be {"" 0 false} in current case
	fmt.Println("The value:", v, "Is Present?", ok)

}

// MapExecuteTask to execute word
func MapExecuteTask() {
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// FunctionValues to understand how functions can be variables, and passed around.
func FunctionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	} // Creating function variables

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
