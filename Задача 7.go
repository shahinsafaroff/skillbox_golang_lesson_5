
//7. Счастливый билет
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Println("Введите первую число: ")
	fmt.Scan(&a)
	fmt.Println("Введите вторую число: ")
	fmt.Scan(&b)
	fmt.Println("Введите третью число: ")
	fmt.Scan(&c)
	fmt.Println("Введите четвертую число: ")
	fmt.Scan(&d)
	if (a+b)%(c+d) == 0 {
		fmt.Println("счастливый билет")
	} else if a == d && b == c {
		fmt.Println("зеркальный билет")
	} else {
		fmt.Println("обычный билет")
	}
}
