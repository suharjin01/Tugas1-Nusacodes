package main

import (
	"fmt"
	"sync"
)

// Kalkulator struct dengan atribut result dan mutex
type Kalkulator struct {
	result int
	mutex  sync.Mutex
}

// Tambah menambahkan nilai ke result secara aman
func (k *Kalkulator) Tambah(n int) {
	k.mutex.Lock()         // Mengunci akses
	defer k.mutex.Unlock() // Membuka akses setelah selesai
	k.result += n
}

// Kurang mengurangi nilai dari result secara aman
func (k *Kalkulator) Kurang(n int) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.result -= n
}

// Kali mengalikan nilai result dengan n secara aman
func (k *Kalkulator) Kali(n int) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.result *= n
}

// Hasil mengembalikan nilai result secara aman
func (k *Kalkulator) Hasil() int {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	return k.result
}

func main() {
	kalk := &Kalkulator{}

	var wg sync.WaitGroup

	// Fungsi untuk menjalankan operasi matematika
	operasi := func(op string, n int) {
		defer wg.Done() // Kurangi jumlah goroutine di WaitGroup setelah selesai
		switch op {
		case "add":
			kalk.Tambah(n)
			fmt.Printf("Menambahkan %d: Hasil sementara = %d\n", n, kalk.Hasil())
		case "subtract":
			kalk.Kurang(n)
			fmt.Printf("Mengurangi %d: Hasil sementara = %d\n", n, kalk.Hasil())
		case "multiply":
			kalk.Kali(n)
			fmt.Printf("Mengalikan %d: Hasil sementara = %d\n", n, kalk.Hasil())
		default:
			fmt.Printf("Operasi %s tidak dikenal\n", op)
		}
	}

	// Menambahkan operasi ke WaitGroup
	wg.Add(3)

	// Jalankan operasi secara paralel menggunakan goroutine
	go operasi("add", 10)     // Menambahkan 10
	go operasi("subtract", 5) // Mengurangi 5
	go operasi("multiply", 3) // Mengalikan dengan 3

	// Tunggu semua operasi selesai
	wg.Wait()

	fmt.Println("Hasil akhir:", kalk.Hasil())
}
