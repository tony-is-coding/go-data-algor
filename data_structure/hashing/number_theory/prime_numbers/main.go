package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

var specialNonPrime = []int{0, 1, 2}

func checkPrime(number int) bool {
	if ok := funk.Contains(specialNonPrime[:2], number); ok {
		return false
	}
	if specialNonPrime[len(specialNonPrime)-1] == number {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(checkPrime(11))
}
