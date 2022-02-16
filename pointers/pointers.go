package main

import "fmt"

func main() {
	// Напишите программу, которая меняет местами два числа (x := 1; y := 2; swap(&x, &y) должно дать x=2 и y=1)

	x := 1
	y := 2

	fmt.Println("addresses ", &x, &y)

	fmt.Println("initial values ", x, y)
	swap(&x, &y)
	fmt.Println("changed values ", x, y)

	valueReplacement(&x, &y)

	fmt.Println(x, y)

	var z = 1.5
	fmt.Println("square z = ", square(&z))
}

func valueReplacement(x *int, y *int) {
	*x = 33
	*y = 25
}

func swap(i *int, i2 *int) {
	tmp := *i
	fmt.Println("tmp = ", tmp)
	*i = *i2
	*i2 = tmp
	fmt.Println("addresses ", i, i2)
}

func square(z *float64) float64 {
	*z = *z * *z
	return *z
}
