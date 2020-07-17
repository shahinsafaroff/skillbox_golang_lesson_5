//1. Определение координатной плоскости точки

package main

import "fmt"

func main() {
	var x, y float64
	fmt.Println("Введите первую число: ")
	fmt.Scan(&x)
	fmt.Println("Введите вторую число: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("точка находится в первой четверти координатной плоскости")
	} else if x < 0 && y > 0{
		fmt.Println("точка находится во второй четверти")
	} else if x < 0 && y < 0 {
		fmt.Println("число находится в третьей четверти")
	} else if x > 0 && y < 0 {
		fmt.Println("точка лежит в четвертой четверти")
	}
}
