package main

import "fmt"

func odd(num int) bool {
	if num%2 == 1 {
		return true
	} else {
		return false
	}
}
func even(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for x := 2; x < num; x++ {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func main() {
	// even number
	fmt.Println("even number :")
	for i := 0; i < 1001; i++ {
		if even(i) {
			fmt.Print(i, " ")
		}
	}
	// odd number
	fmt.Println("\nodd  number :")
	for i := 0; i < 1001; i++ {
		if odd(i) {
			fmt.Print(i, " ")
		}
	}
	//multiplies by 5
	fmt.Println("\nmultipies by 5:")
	for i := 0; i < 1001; i++ {
		fmt.Print(i*5, " ")
	}

	//prime numbers
	fmt.Println("\nprime  number :")
	for i := 0; i < 1001; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	//prime numbers
	// fmt.Println("\nprime  number less than 100 :")
	// for i := 0; i < 101; i++ {
	// 	if isPrime(i) {
	// 		fmt.Print(i, " ")
	// 	}
	// }

}
