//3. Проверить, что есть совпадающие числа

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите первую число: ")
	fmt.Scan(&a)
	fmt.Println("Введите вторую число: ")
	fmt.Scan(&b)
	fmt.Println("Введите третью число: ")
	fmt.Scan(&c)

	if a == b || b == c || a == c {
		fmt.Println("два или более числа совпадают")
	} else {
		fmt.Println("ни один из чисел не совпадают")
	}
}