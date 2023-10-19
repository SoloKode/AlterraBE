package main

import "fmt"

func main() {
	var pilihan uint8
	fmt.Println("1. Bilangan Prima")
	fmt.Println("2. Bilangan Kelipatan Tujuh")
	fmt.Println("3. Hitung Luas Trapesium")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilihan)
	if pilihan <= 1 {
		bilanganPrima()
	} else if pilihan == 2 {
		kelipatanTujuh()
	} else {
		luasTrapesium()
	}

}

func bilanganPrima() {
	var bilangan int
	fmt.Print("Input Bilangan : ")
	fmt.Scanln(&bilangan)
	fmt.Println("Bilangan Input : ", bilangan)
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
func kelipatanTujuh() {
	var bilangan int
	fmt.Print("Input Bilangan : ")
	fmt.Scanln(&bilangan)
	fmt.Println("Bilangan Input : ", bilangan)
	kelipatan := 7

	if bilangan%kelipatan == 0 {
		fmt.Println("Bilangan Kelipatan Tujuh")
	} else {
		fmt.Println("Bukan Bilangan Kelipatan Tujuh")
	}
}
func luasTrapesium() {
	var t float32
	var A float32
	var B float32
	var L float32
	fmt.Print("Masukkan Nilai t: ")
	fmt.Scanln(&t)
	fmt.Print("Masukkan Nilai A: ")
	fmt.Scanln(&A)
	fmt.Print("Masukkan Nilai B: ")
	fmt.Scanln(&B)
	L = 0.5 * t * (A + B)
	fmt.Println("Luas Trapesium : ", L)
}
