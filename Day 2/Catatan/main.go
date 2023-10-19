package main

import "fmt"

func main() {
	n := 5
	for j := 1; j <= n; j++ {
		for k := n - 1; k >= j; k-- {
			fmt.Print(" ")
		}
		for i := 1; i <= j; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	// for j := 1; j <= 5; j++ {
	// 	for i := 1; i <= j; i++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	//bitwise >> << (pergeseran bit)
	// a := 2 // 10 (biner)
	// b := 2
	// hasil := a << b // 100 (biner)
	// fmt.Println(hasil)
	// if age > 17 {
	// 	fmt.Println("Dewasa")
	// } else if age > 10 {
	// 	fmt.Println("Anak-anak")
	// } else {
	// 	fmt.Println("Invalid")
	// }
	// age := 17
	// switch {
	// case age > 17:
	// 	fmt.Println("Dewasa")
	// case age > 10:
	// 	fmt.Println("Anak-anak")
	// default:
	// 	fmt.Println("Invalid")
	// }
}
