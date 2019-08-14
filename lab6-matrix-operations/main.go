package main

import "fmt"

// solution 1
func main() {

	matrixS := [4][4]interface{}{
		{"a", 4, 5, 0},
		{"d", 0, 3, 0},
		{"m", 4, 8, 0},
		{"s", 1, 4, 0},
	}

	functions := map[string]func(int, int) int{
		"a": func(a int, b int) int { return a + b },
		"d": func(a int, b int) int { return a / b },
		"m": func(a int, b int) int { return a * b },
		"s": func(a int, b int) int { return a - b },
	}

	for index, row := range matrixS {
		fName := ((row[0]).(string))
		l1 := row[1].(int)
		r1 := row[2].(int)
		matrixS[index][3] = functions[fName](l1, r1)
	}

	fmt.Println(matrixS)
}

// solution 2
// func main() {

// 	a := func(a int, b int) int { return a + b }
// 	s := func(a int, b int) int { return a - b }
// 	m := func(a int, b int) int { return a * b }
// 	d := func(a int, b int) int { return a / b }

// 	matrixF := [4][4]interface{}{
// 		{a, 4, 5, 0},
// 		{d, 0, 3, 0},
// 		{m, 4, 8, 0},
// 		{s, 1, 4, 0},
// 	}

// 	for index, row := range matrixF {
// 		f := row[0].(func(int, int) int)
// 		l1 := row[1].(int)
// 		r1 := row[2].(int)
// 		matrixF[index][3] = f(l1, r1)
// 	}

// 	fmt.Println(matrixF)
// }
