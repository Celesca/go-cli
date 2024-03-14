package main

import "fmt"

func main() {

	var MyArray [5]int
	MyArray[0] = 1
	MyArray[1] = 2
	MyArray[2] = 3
	MyArray[3] = 4
	MyArray[4] = 5

	fmt.Println("Array: ", MyArray)

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice: ", mySlice)

	mySlice = append(mySlice, 6)
	fmt.Println("Slice after append: ", mySlice)

	mySlice = append(mySlice[:3], mySlice[4:]...)
	fmt.Println("Slice after delete: ", mySlice)

	subSlice := mySlice[1:3]
	fmt.Println("Subslice: ", subSlice)

}
