package main

import(
	"fmt"
)

func Mergesort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	pivot := len(items)/2

	return Merge(Mergesort(items[:pivot]), 
		Mergesort(items[pivot:]) )
}

func Merge(left, right []int) []int {
	var output []int

	for len(left)!=0 && len(right)!=0 {
		if left[0] < right[0] {
			output = append(output, left[0])
			left = left[1:]
		} else {
			output = append(output, right[0])
			right = right[1:]	
		}
	}
	if len(left) >0 {
		output = append(output, left...)
	} else if len(right)>0 {
		output = append(output, right...)
	}
	return output
}

func main() {
	var array []int = []int {5, 4, 3, 2, 1, 0, 10}	
	array = Mergesort(array)
	fmt.Println(array)
}