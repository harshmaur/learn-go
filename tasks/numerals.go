package tasks

import "fmt"

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
