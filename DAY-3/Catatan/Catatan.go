package main

import (
	"fmt"
	"slices"
)

func main() {
	// // add
	// add(10, 20)
	// // addReturn
	// result := addReturn(40, 20)
	// fmt.Println("Hasil : ", result)
	// //multiReturnValue
	// hasila, hasilb := multiReturnValue(99, 66)
	// fmt.Println("Hasil A : ", hasila)
	// fmt.Println("Hasil B : ", hasilb)
	// strukturDataArray()
	// strukturDataSlice()
	strukturMap()
}
func add(a int, b int) {
	fmt.Println("Hasil : ", a+b)
}
func addReturn(a int, b int) int {
	return a + b
}
func multiReturnValue(a int, b int) (int, int) {
	return a, b
}
func strukturDataArray() {
	//Array, Banyak Data, 1 Tipe, Tidak Bisa Dikurangi Atau Ditambah
	// var scoreData [5]int = [5]int{100, 90, 80, 7}
	scoreData := [5]int{100, 90, 80, 7}
	scoreData[2] = 200

	// for _, v := range scoreData {
	// 	fmt.Println(v)
	// }
	for i, v := range scoreData {
		fmt.Println("Nilai score ke- ", i, " adalah ", v)
	}
	// fmt.Println(scoreData)
	// for i := 0; i < len(scoreData); i++ {
	// 	fmt.Println(i, scoreData[i])
	// }
}
func strukturDataSlice() {
	// var score []int
	// var score []int = []int{1, 2, 3}
	score := []int{1, 2, 3, 4, 5}
	slices.Delete(score, 2, 3)
	fmt.Println(score)
	// score = append(score, 10)
	// score[0] = 20
	// fmt.Println(score[0:3])
	// fmt.Println(score[:3])
	// fmt.Println(score[3:])
	// scoreNew := make([]int, 100)
	// fmt.Print(scoreNew)
}
func strukturMap() {
	// Januari : 10k
	// Februari : 20k
	// Maret : 30k
	// map => key, value => key unique

	// dataPendapatan := make(map[string]int)
	dataPendapatan := map[string]int{}
	dataPendapatan["Januari"] = 1000
	dataPendapatan["Februari"] = 2000
	dataPendapatan["Maret"] = 3000
	delete(dataPendapatan, "Januari")
	// result, isFound := dataPendapatan["Februari"]
	result, isFound := dataPendapatan["Februarissaassaa"]
	// fmt.Println(dataPendapatan)
	fmt.Println(result, isFound)
	fmt.Println(dataPendapatan)
}
