//8. Игра “Угадывание числа” (дополнительно)
package main

import "fmt"

func main() {
	fmt.Println("Keep the number from 1 to 10 and I will find your number")

	min := 1
	max := 10
	i := 0
	for i < 11 {
		number := (min + max) / 2
		fmt.Println("Founded number:", number, "? type '=' if your number was found, '>' if your number bigger, '<' if your number lesser)")

		var answer string
		fmt.Scan(&answer)
		if answer == ">" {
			min = number + 1
		} else if answer == "<" {
			max = number - 1
		} else if answer == "=" {
			return
		} else {
			fmt.Println("HOOORAAH! Number is found")
		}

	}
}
