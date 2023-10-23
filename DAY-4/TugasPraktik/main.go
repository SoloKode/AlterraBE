package main

import "fmt"

// Struct Mobil
type Car struct {
	Name   string
	Color  string
	FuelIn float32
}

func (car Car) estimasi() float32 {
	return car.FuelIn / 1.5
}
func Soal1() {
	var car1 Car = Car{"Volvo", "Hitam", 2.50}
	fmt.Printf("Dengan Bahan Bakar : %.2f Liter\n", car1.FuelIn)
	fmt.Printf("Estimasi Jarak : %.3f", car1.estimasi())
}

// Struct Student
type Student struct {
	Name  []string
	Score []int
}

func (students Student) rataRata() (min, max, rata, maxIndex, minIndex int) {
	max = 0
	min = 100
	var tempNilai int = 0
	for i, nilai := range students.Score {
		if nilai > max {
			max = nilai
			maxIndex = i
		}
		if nilai < min {
			min = nilai
			minIndex = i
		}
		tempNilai += nilai

	}
	rata = tempNilai / len(students.Score)
	return min, max, rata, maxIndex, minIndex
}
func Soal2() {
	var students Student = Student{Name: []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"}, Score: []int{80, 75, 70, 75, 60}}
	fmt.Println(students)
	min, max, rata, maxIndex, minIndex := students.rataRata()
	fmt.Println("Average Score : ", rata)
	fmt.Printf("Min Score of Students: %s (%d)\n", students.Name[minIndex], min)
	fmt.Printf("Min Score of Students: %s (%d)\n", students.Name[maxIndex], max)
}
func getMinMax(numbers ...*int) (min int, max int) {
	max = *numbers[0]
	min = *numbers[0]
	for _, value := range numbers {
		if *value > max {
			max = *value
		}
		if *value < min {
			min = *value
		}
	}
	return min, max
}
func Soal3() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&a3)
	fmt.Scanln(&a4)
	fmt.Scanln(&a5)
	fmt.Scanln(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min : ", min)
	fmt.Println("Nilai max : ", max)

}
func main() {
	var pilihan uint8
	fmt.Println("1. Tugas Method Perkiraan Jarak")
	fmt.Println("2. Tugas Struct Students")
	fmt.Println("3. Tugas Nilai Min & Max")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilihan)
	if pilihan <= 1 {
		Soal1()
	} else if pilihan == 2 {
		Soal2()
	} else {
		Soal3()
	}
}
