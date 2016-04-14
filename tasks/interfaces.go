package tasks

import (
	"fmt"
	"math"
	"sort"
)

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	area() float64
}

func ShapeOperations(s Shape) {
	fmt.Println(s.area())
}

func RunShape() {
	s := Square{10}
	c := Circle{6}
	ShapeOperations(s)
	ShapeOperations(c)
}

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func SortPeople() {
	p := people{"Harsh", "John", "Zeno", "Al", "Todd"}
	sort.Sort(p)
	fmt.Println(p)

}

func SortPeopleReverse() {
	p := people{"Harsh", "John", "Zeno", "Al", "Todd"}
	sort.Sort(sort.Reverse(p))
	fmt.Println(p)
}

func SortStringSlice() {
	p := []string{"Harsh", "John", "Zeno", "Al", "Todd"}
	sort.Sort(sort.StringSlice(p))
	fmt.Println(p)
}

func SortIntSlice() {
	p := []int{2, 1, 5, 4, 6, 2, 4}
	sort.Sort(sort.IntSlice(p))
	fmt.Println(p)
}
