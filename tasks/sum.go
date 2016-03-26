package tasks

import (
	"fmt"
	"reflect"
)

// Zerosum is to find three numbers whose sum is zero given an integer slice
func Zerosum(someSlice []int) {
	newSlice := make([]int, 3)
	for i1, v1 := range someSlice {
		x := v1
		for i2, v2 := range someSlice {
			if i1 != i2 {
				y := v2
				for i3, v3 := range someSlice {
					if i1 != i2 && i1 != i3 && i2 != i3 {
						z := v3
						if x+y+z == 0 {
							newSlice2 := []int{x, y, z}
							if !reflect.DeepEqual(newSlice, newSlice2) {
								newSlice = []int{x, y, z}
								fmt.Println(newSlice)
							}
						}
					}
				}
			}
		}
	}

}
