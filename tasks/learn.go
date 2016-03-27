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
