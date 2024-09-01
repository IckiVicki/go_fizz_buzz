package main

import "fmt"

// This is a comment

var fizz int8 = 3
var buzz int8 = 5

var number_fizzing int16 = 50

func main() {
	for i := 0; i <= int(number_fizzing); i++ {
		if i != 0 {
			if i%int(buzz) == 0 {
				if i%int(fizz) == 0 {
					fmt.Println("Fizzbuzz")
				}
				if i%int(fizz) != 0 {
					fmt.Println("Buzz")
				}
			}
			if i%int(fizz) == 0 {
				fmt.Println("Fizz")
			}
			if i%int(fizz) != 0 {
				if i%int(buzz) != 0 {
					fmt.Println(i)
				}
			}
		}
	}
}
