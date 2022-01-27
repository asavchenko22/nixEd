package main

import (
	"fmt"
	"strconv"
)

func main() {
	flowControl(100)
}

func flowControl(a int) {
	for i := 0; i <= 100; i++ {
		result := strconv.Itoa(a)
		switch a % 3 {
		case 0:
			result = "fizz"
		}
		switch a % 5 {
		case 0:
			result = "buzz"
		}
		if (a%3 == 0) && (a%5 == 0) {
			result = "fizzBuzz"
		}
		fmt.Println(result)
		a += 1
	}
}
