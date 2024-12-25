package main

import "fmt"

// Fungsi Add
func add(value1, value2 int) int {
	jumlah := value1 + value2
	return jumlah
}

// Fungsi Bagi
func bagi(angka1, angka2 int) int {
	bagi := angka1 / angka2
	return bagi
}

// Fungsi Variadic
func kurangVariadic(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total -= number
	}

	return total
}

func main() {
	// memanggil fungsi add()
	hasil1 := add(50, 50)
	fmt.Println("Hasil dari fungsi add() adalah : ", hasil1)

	// memanggil fungsi bagi()
	hasil2 := bagi(100, 50)
	fmt.Println("Hasil dari fungsi bagi() adalah : ", hasil2)

	// memanggil fungsi kurangVariadic()
	hasil3 := kurangVariadic(100, 20, 10, 5, 15)
	fmt.Println("Hasil dari fungsi kurangVariadic() adalah : ", hasil3)
}
