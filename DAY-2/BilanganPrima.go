package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scanln(&bilangan)
	fmt.Println(bilangan)
	A := 2
	if bilangan <= 1 {
		fmt.Println("Bukan Bilangan Prima")
	} else {
		for A < bilangan {
			if bilangan%A == 0 {
				fmt.Println("Bukan Bilangan Prima")
				break
			}
			A++
		}
		if A == bilangan {
			fmt.Println("Bilangan Prima")
		}
	}
}
