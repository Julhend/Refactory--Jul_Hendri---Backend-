package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{3, 12, 4, 5, 8, 9}

	fmt.Println("Before:")
	fmt.Println("Number: ", num)
	sort.Ints(num)
	fmt.Println("\nSorted:")
	fmt.Println("Number : ", num)

}
