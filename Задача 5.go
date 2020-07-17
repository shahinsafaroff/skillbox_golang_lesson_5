//5. Определение максимальных процентов

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите первую процентные ставку: ")
	fmt.Scan(&a)
	fmt.Println("Введите вторую процентные ставку: ")
	fmt.Scan(&b)
	fmt.Println("Введите третью процентные ставку: ")
	fmt.Scan(&c)
	if a > c && b > c  {
		fmt.Println("первая и вторая  более выгодные ставки")
	} else if a > b && c > b {
		fmt.Println("первая и третья  более выгодные ставки")
	} else if b > c && a > c  {
		fmt.Println("вторая и первая  более выгодные ставки")
	} else if c > a && b > a {
		fmt.Println("третья и вторая  более выгодные ставки")
	} else {
		fmt.Println("Два или болще ставки равны!!!")
	}
}