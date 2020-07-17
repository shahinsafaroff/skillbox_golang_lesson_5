//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Enter your income: ")
//	var profit, tax float64
//	fmt.Scan(&profit)
//
//
//
//	if profit > 50000 {
//		tax = (profit - 50000) * 30 / 100 + ((50000 - 10000) * 20 / 100) + (10000 * 13 /100)
//		fmt.Println("Proportion of tax is equal: ", tax)
//	} else if profit > 10000 && profit < 50000 {
//		tax = (((profit - 10000) * 20 / 100) + (10000 * 13 /100))
//		fmt.Println("Proportion of tax is equal: ", tax)
//	} else if profit > 0 && profit < 10000 {
//		tax = profit * 13 / 100
//		fmt.Println("Proportion of tax is equal: ", tax)
//	} else if profit == 0 {
//		fmt.Println("Really???")
//	} else {
//		fmt.Println("You've mistyped")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Enter your income: ")
//	var income, tax float64
//	fmt.Scan(&income)
//
//	if income > 10000 {
//		tax = 10000 * 0.13
//		if income > 50000 {
//			tax += (50000-10000) * 0.2 + (income-50000) * 0.3
//		} else if income < 50000 && income > 10000 {
//			tax += (income - 10000) * 0.2
//		} else {
//			tax = 10000 * 0.13
//		}
//		fmt.Println("Proportion of tax is equal: ", tax)
//	} else if income <= 0 {
//		fmt.Println("Really???")
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var name string
//	fmt.Println("Enter the Name: ")
//	fmt.Scan(&name)
//
//	var isUser bool
//	isUser = true
//
//	isAdmin := name == "admin"
//
//	fmt.Println("entered name is user", isUser)
//	fmt.Println("entered name is admin", isAdmin)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var a, b, c int
//	fmt.Println("Введите первую число: ")
//	fmt.Scan(&a)
//	fmt.Println("Введите вторую число: ")
//	fmt.Scan(&b)
//	fmt.Println("Введите третью число: ")
//	fmt.Scan(&c)
//
//	if a > b && a > c {
//	fmt.Println(a)
//	} else if b > c {
//		fmt.Println(b)
//	} else {
//		fmt.Println(c)
//	}
//}
