package tasks

import "fmt"

const (
	_  = iota             // iota = 0
	KB = 1 << (iota * 10) // shifts 1 to the left 10 places // iota = 1
	MB = 1 << (iota * 10) // iota = 2
	GB = 1 << (iota * 10) // 3
	TB = 1 << (iota * 10) // 4

	metersToYards = 1.09361
)

// Print42 prints 42
func Print42() {
	fmt.Println(42)
}

func Print42Binary() {
	fmt.Printf("%d - %b", 42, 42)
}

func Print42Formats() {
	// fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	// fmt.Printf("%d - %b - %#x\n", 42, 42, 42)

	for i := 1; i < 100; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

func BitWiseShifting() {
	fmt.Printf("%d - %b\n", KB, KB)
	fmt.Printf("%d - %b\n", MB, MB)
	fmt.Printf("%d - %b\n", GB, GB)
	fmt.Printf("%d - %b\n", TB, TB)
}

func MemoryAddress() {
	a := 42
	fmt.Println("a - ", a)
	fmt.Println("a address", &a)
	// fmt.Printf("%d\n", &a)

	b := &a
	fmt.Println("b is", b)
	fmt.Println("b dereferenced is", *b)

	*b = 43
	fmt.Println("a is now ", a)
	// var meters float64
	// fmt.Print("Enter Meters:")
	// fmt.Scan(&meters) // You need to add reference/opeartor to make sure it works
	// yards := meters * metersToYards
	// fmt.Println(meters, " meter is ", yards, " yards")
}

func somePointer(z *int) {
	fmt.Println(z)
	*z = 0
}

func PointersHelp() {
	x := 5
	fmt.Println(&x)
	somePointer(&x)
	fmt.Println(x)
}

func NewForLoop() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func NestedForLoop() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}

func Runes() {

	// utf-8 default 1-4 byte scheme.
	// rune is an alias of int32 or 4 Bytes.
	// foo = 'a' // notice single quote to tell it is rune
	for i := 500; i < 5000; i++ {
		fmt.Println(string(i), []byte(string(i)))
	}
}

func Strings() {
	// strings can be identified by backticks `` or double quotes ""
}

func Switch() {
	// Switch without an expression can be used to write multiple if else's
	// fallthough will make the next case true in all cases.

}

// VariadicAverage ... allows to accept 0 - many variables
// type after ... specifies what type of slice of variables it is going to be. so in this case []float64
func VariadicAverage(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
	// Call through the following ways

	// tasks.VariadicAverage(10, 15, 20, 25, 30)

	// data := []float64{10, 15, 20, 25, 30}
	// tasks.VariadicAverage(data...) the three dots at end will split the slice and pass as arguments.
}

func makeGreeter() func() string {
	return func() string {
		return "Hello World!"
	}
}

func Closure1() {
	greet := makeGreeter()
	fmt.Println(greet())
}

// Callback Example
func filter(numbers []int, callback func(a int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}
func Call1() {
	xs := filter([]int{1, 2, 3, 4}, func(a int) bool {
		return a > 1
	})
	fmt.Println(xs)
}

// closure happens normally when function is RETURNED
// Callbacks happens normally when the function is a passed as argument
//Recursion happens when a function is called in itself.

func RecursionFactorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * RecursionFactorial(x-1)
}

func changeMe(aslice []string) {
	aslice[0] = "harsh"
	fmt.Println(aslice)
}
func Slice1() {
	m := make([]string, 1, 15)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func MultiDSlice() {
	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 3)
		for j := 0; j < 3; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}

func Ascii() {
	for i := 500; i < 2000; i++ {
		fmt.Printf("%d - ", i)
		fmt.Println(string(i))
	}
}
