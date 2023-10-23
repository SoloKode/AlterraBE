package main

import (
	"fmt"
	"strconv"
)

func TugasMergeArray() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// Output yang diharapkan: ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// Output yang diharapkan: ["sergei","jin","steve","bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// Output yang diharapkan: ["alisa","yoshimitsu","devil","jin","law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// Output yang diharapkan: ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// Output yang diharapkan: ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// Output yang diharapkan: []
}
func ArrayMerge(arrayA, arrayB []string) []string {
	mergeArray := make([]string, 0)
	mergeArray = append(mergeArray, arrayA...)

	for _, B := range arrayB {
		found := false
		for _, A := range mergeArray {
			if A == B {
				found = true
				break
			}
		}
		if !found {
			mergeArray = append(mergeArray, B)
		}
	}
	return mergeArray
}
func main() {

	var pilihan uint8
	fmt.Println("1. Tugas Merge Array")
	fmt.Println("2. Tugas Slice Map")
	fmt.Println("2. Tugas Bilangan Muncul Sekali")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scanln(&pilihan)
	if pilihan <= 1 {
		TugasMergeArray()
	} else if pilihan == 2 {
		TugasMappingSlice()
	} else {
		TugasMunculSekali()
	}
}
func Mapping(slice []string) map[string]int {
	dataSlice := map[string]int{}
	for i, data := range slice {
		_, ketemu := dataSlice[data]
		if !ketemu {
			dataSlice[data] = i
		}
	}
	return dataSlice
}
func TugasMappingSlice() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[asd:0 qwe:1 adi:3]
	fmt.Println(Mapping([]string{}))
}

func munculSekali(angka string) []int {
	muncul := make(map[int]int)
	for _, digit := range angka {
		digitInt := int(digit - '0')
		muncul[digitInt]++
	}
	sekali := []int{}
	for _, digit := range angka {
		digitInt := int(digit - '0')
		if muncul[digitInt] == 1 {
			sekali = append(sekali, digitInt)
		}
	}

	return sekali
}

func TugasMunculSekali() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}

func appearOnce(number string) []int {
	mapResult := map[int]int{}
	result := []int{}
	for _, val := range number {
		res, _ := strconv.Atoi(string(val))
		mapResult[res]++
	}
	for k, v := range mapResult {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}

// func Dinamis(slice []string) map[string]int {
// 	var n int
// 	fmt.Scanln(&n)
// 	slice := make([]string, n)
// 	for i,_ := range slice {
// 		fmt.Scanln(&slice[i])
// 	}
// }

// func ArrayMergeMentor(sliceA,sliceB []string) []string {
// 	result := []string{}
// 	mapResult := make(map[string]bool)

// 	result = append(sliceA,sliceB...)
// 	for _, value := range sliceA {
// 		_, isFound := mapResult[value]
// 		if !isFound {
// 			result = append(result,value)
// 			mapResult[value] = true
// 		}
// 	}
// 	return result
// }
