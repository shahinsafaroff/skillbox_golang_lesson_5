////6. Решение квадратного уравнения
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, D float64
	fmt.Println("квадратные уравнения: ")
	fmt.Println("Введите первую число: ")
	fmt.Scan(&a)
	fmt.Println("Введите вторую число: ")
	fmt.Scan(&b)
	fmt.Println("Введите третью число: ")
	fmt.Scan(&c)

	D = math.Pow(b,2) - 4 * a * c


	fmt.Println(D)

	if D > 0 {
		X1 := (-b + math.Sqrt(D)) / 2 * a
		X2 := (-b - math.Sqrt(D)) / 2 * a
		fmt.Println(" X1 != X2","дискриминант", D, "корни уравнения", X1, X2)
	} else if D == 0 {
		X := (-b)/ 2 * a
		fmt.Println(" X1 = X2", "корень", X)
	} else {
		fmt.Println("корней нет")

	}
}