package tasks

import "fmt"

// walk attribute for the person
// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int).
func (p Person) walk() {
	fmt.Println("Walked One step")
}

func (p *Person) changeName(name string) {
	p.Name = name
}

// MakePersonWalk a step
func MakePersonWalk() {
	p := Person{"Harsh", 21, true}
	p.walk()
}

// ChangePersonName to hey
func ChangePersonName() {
	p := Person{"Harsh", 21, true}
	fmt.Println(p)
	p.changeName("Hey")
	fmt.Println(p)
}
