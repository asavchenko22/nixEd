package main

import "fmt"

// func makeEvenGenerator() func() int {
// 	i := 0
// 	return func() (ret int) {
// 		ret = i
// 		i += 2
// 		return
// 	}
// }

func factorial(x int) int {

	if x == 0 {
		return 1
	}
	fmt.Println("x = ", x)
	fmt.Println("x*x-1 = ", x*x-1)
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
	// nextEven := makeEvenGenerator()
	// fmt.Println(nextEven()) // 0
	// fmt.Println(nextEven()) // 2
	// fmt.Println(nextEven()) // 4
}
