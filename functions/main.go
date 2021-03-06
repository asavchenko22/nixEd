package main

import (
	"fmt"
	//"strconv"
)

func main() {
	makeSlise()

	fmt.Println(evenOrOdd(2))

	fmt.Println(findGreatNumber())

	nextOdd := makeOddGenerator()
	fmt.Println("nextOdd : ", nextOdd())
	fmt.Println("nextOdd : ", nextOdd())
	fmt.Println("nextOdd : ", nextOdd())

	// Последовательность чисел Фибоначчи определяется как fib(0) = 0, fib(1) = 1,
	// fib(n) = fib(n-1) + fib(n-2). Напишите рекурсивную функцию, находящую fib(n).

	fmt.Println("fib(7) = ", fib(5, true))
}

func fib(i int, print bool) int {
	if print {
		n := 0
		nPrev := 0
		nPrevPrev := 0
		for j := 0; j <= i; j++ {
			switch j {
			case 0:
				fmt.Println(n)
				nPrev++
				nPrevPrev++
			case 1:
				n = nPrev
				fmt.Println(n)
			case 2:
				fmt.Println(n)
				n = nPrev + nPrevPrev
			default:
				fmt.Println(n)
				nPrevPrev = nPrev
				nPrev = n
				n = nPrevPrev + nPrev

			}
		}

	}
	print = false
	if i > 1 {
		i = fib(i-1, print) + fib(i-2, print)
	}
	return i
}

func makeOddGenerator() func() int {
	i := int(1)
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func findGreatNumber() int {
	res := 0
	s := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	for _, v := range s {
		if res < v {
			res = v
		}
	}

	return res
}

func evenOrOdd(x int) (int, bool) {
	res := x / 2
	bol := false
	if x%2 == 0 {
		bol = true
	}
	return res, bol
}

func makeSlise() {
	a := make([]int, 5)
	for i := range a {
		fmt.Println("i = ", i)
		a[i] = i
	}
	fmt.Println("sum Slise : ", sum(a))
}

func sum(x []int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}
