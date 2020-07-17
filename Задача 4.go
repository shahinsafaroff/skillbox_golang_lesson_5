// 4. Сумма без сдачи

package main

import "fmt"

func main() {
	var summ, a, b, c int
	fmt.Println("Сумма за товар: ")
	fmt.Scan(&summ)
	fmt.Println("Введите первую число: ")
	fmt.Scan(&a)
	fmt.Println("Введите вторую число: ")
	fmt.Scan(&b)
	fmt.Println("Введите третью число: ")
	fmt.Scan(&c)

	if summ == a + b + c || summ == a + b || summ == b + c {
		fmt.Println("Пользователь может оплатить без сдачи")
	} else {
		fmt.Println("Пользователь не может оплатить без сдачи")

	}
}