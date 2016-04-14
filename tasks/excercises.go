package tasks

import "fmt"

func HelloWorld() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}

func SmallLarge() {
	var num1, num2 int
	fmt.Print("Enter Smaller Number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter Larger Number: ")
	fmt.Scan(&num2)
	fmt.Println(num2 % num1)
}

func PrintEven() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func FizzBuzz() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func SumDivisible() {
	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i

		}
	}
	fmt.Println(sum)
}

func EvenDivideTwo(n int) (int, bool) {
	if n%2 == 0 {
		return (n / 2), true
	}
	return (n / 2), false
}

func EvenDivideTwoExp() {
	divide := func(n int) (int, bool) {
		if n%2 == 0 {
			return (n / 2), true
		}
		return (n / 2), false
	}
	fmt.Println(divide(4))
}

func GreatestNumber(ns ...int) int {
	var large int
	for _, v := range ns {
		if v > large {
			large = v
		}
	}
	return large
}
