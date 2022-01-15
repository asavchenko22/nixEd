package main

import (
	"fmt"
)

func intermediaRating(x [3]float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

func testIMR(value ...float64) float64 {
	// var x = make([]float64, len(value))
	var x [3]float64
	for i, v := range value {
		x[i] = v
		i++
	}
	return intermediaRating(x)

	// ----------------------------------

	// var total float64
	// for _, v := range value {
	// 	total += v
	// }
	// return total / float64(len(value))
}

// func printInt(a int) {
// 	for i := 0; i <= 100; i++ {
// 		result := strconv.Itoa(a)
// 		switch a % 3 {
// 		case 0:
// 			result = "fizz"
// 		}
// 		switch a % 5 {
// 		case 0:
// 			result = "buzz"
// 		}
// 		if (a%3 == 0) && (a%5 == 0) {
// 			result = "fizzBuzz"
// 		}
// 		fmt.Println(result)
// 		a += 1
// 	}
// }

func main() {
	fmt.Println("Hello, NIX Education")
	fmt.Println(testIMR(3.0, 3.0, 3.0))

	fmt.Println("S--------L--------I---------S--------E")

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	fmt.Println("len(slise1) = ", len(slice1))
	fmt.Println("len(slise2) = ", len(slice2))

	fmt.Println("M-----------------A------------------P")

	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println("questions--------------------answers")

	var y = make([]int, 3, 9)
	fmt.Println("z len = ", len(y))
	fmt.Println("z cap = ", cap(y))

	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("x[2:5] = ", x[2:5])

	xy := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println("findLess = ", findLess(xy))

	//printInt(0)
}

func findLess(xy []int) int {
	res := xy[0]
	for _, v := range xy {
		if res > v {
			res = v
		}
	}
	// fmt.Println("xy[15] = ", xy[15])
	// for i := 1; i < len(xy); i++ {
	// 	fmt.Println("i = ", i)
	// 	if res > xy[i] {
	// 		res = xy[i]
	// 	}
	// }
	return res
}
