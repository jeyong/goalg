package main

import(
	"fmt"
)

func Qsort(list []int) []int{
	if len(list) < 2 {
		return list
	}
	low :=0
	pivot := len(list)/2

	list[low], list[pivot] = list[pivot], list[low]
	stack := 0
	for i:=1; i<len(list); i++ {
		if list[i] < list[low] {
			stack++
			list[i], list[stack] = list[stack], list[i]
		}
	}
	list[0], list[stack] = list[stack], list[0]

	var out []int
	out = append(out, Qsort(list[:stack])...)
	out = append(out, list[stack])
	out = append(out, Qsort(list[stack+1:])...)

	return out 
}


func main() {

	array := []int {9, 8, 5, 3, 2, 1, 0}
	Qsort(array)
	fmt.Println("output: ", array)
}