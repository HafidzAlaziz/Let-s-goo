package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func h1() {
	// Nilai Yang Akan Masuk Kedalam Himpunan //
	rand.Seed(time.Now().UnixNano())
	N := 140
	A := make([]int, 3)
	for i := range A {
		A[i] = rand.Intn(N)
	}
	B := make([]int, 3)
	for i := range B {
		B[i] = rand.Intn(N)
	}
	C := make([]int, 2)
	for i := range C {
		C[i] = rand.Intn(N)
	}

	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                      Materi 1                          ║")
	fmt.Println("║                     Himpunan 1                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Printf("Generating Sets... (N=%d)\n", N)
	fmt.Printf("A: %v | B: %v | C: %v\n", A, B, C)

	// Elemen unik yang tidak ada di A atau B //
	angkaHantuAB := Selisihgabungan(A, B)

	// Hapus Elemen C //
	eliminasiC := Selisih(angkaHantuAB, C)

	// Elemen yang ada di A sekaligus di C //
	angkaAC := Irisan(A, C)

	// Gabungan kedua hasil sebelumnya //
	Hasilhimpunan := Gabungan(eliminasiC, angkaAC)
	sort.Ints(Hasilhimpunan) // Mengurutkan hasil agar rapi sesuai output

	fmt.Printf("Hasil Operasi himpunan` : %v\n", Hasilhimpunan)

	// Filter Hasil Himpunan //
	limit := N / 4
	var filtered []int
	for _, val := range Hasilhimpunan {
		if val%2 == 0 && val < limit {
			filtered = append(filtered, val)
		}
	}

	fmt.Printf("Hasil Filter (Genap < %d): %v\n", limit, filtered)
	fmt.Printf("Total Elemen: %d\n", len(filtered))
}

// Cek apakah suatu nilai ada dalam slice
func pengecekan(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Irisan (A n B): Ada di A DAN ada di B
func Irisan(a, b []int) []int {
	var res []int
	for _, v := range a {
		if pengecekan(b, v) {
			res = append(res, v)
		}
	}
	return res
}

// Selisih (A \ B): Ada di A tapi TIDAK ada di B
func Selisih(a, b []int) []int {
	var res []int
	for _, v := range a {
		if !pengecekan(b, v) {
			res = append(res, v)
		}
	}
	return res
}

// Gabungan (A u B): Ada di A atau B
func Gabungan(a, b []int) []int {
	res := a
	for _, v := range b {
		if !pengecekan(res, v) {
			res = append(res, v)
		}
	}
	return res
}

// Selisih Gabungan (A & B): (A \ B) ∪ (B \ A)
func Selisihgabungan(a, b []int) []int {
	diff1 := Selisih(a, b)
	diff2 := Selisih(b, a)
	return Gabungan(diff1, diff2)
}

// -------------------- PEMISAH: PENGELOMPOKAN HP --------------------
// Fungsi: Pengelompokan (dipindahkan dari pengelompokanhp.go)
func h2() {
	rand.Seed(time.Now().UnixNano())
	K := 4
	S := []int{1, 2, 3, 4}
	count := 0

	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                      Materi 1                          ║")
	fmt.Println("║                     Himpunan 2                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Printf("Set S: %v | Target K: %d\n", S, K)
	fmt.Println("Subset 2-Elemen (Sum > 4):")

	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			x := S[i]
			y := S[j]
			sum := x + y

			// Syarat hasil
			if sum > K {
				count++
				fmt.Printf("%d.  {%d, %d} (Sum=%d)\n", count, x, y, sum)
			}
		}
	}

	// Hasil akhir
	fmt.Printf("... Total: %d Pasangan\n", count)
}
