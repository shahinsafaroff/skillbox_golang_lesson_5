//2. Проверить, что одно из чисел положительное

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

	if a > 0 || b > 0 || c > 0 {
		fmt.Println("один из чисел является положительным")
	} else {
		fmt.Println("ни один из чисел не является положительным")
	}
}