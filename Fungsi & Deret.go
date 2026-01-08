package main

import (
	"fmt"
	"math"
)

// Soal 6: Analisis Kedekatan Deret Geometri
func fd2() {
	//
	// Paket No. 8
	a := 9.0
	r := 0.1
	N := 14

	// Hitung jumlah N suku berhingga
	Sn := a * (1 - math.Pow(r, float64(N))) / (1 - r)

	// Hitung jumlah tak hingga
	Sinf := a / (1 - r)

	// Hitung persentase kedekatan
	kedekatan := (Sn / Sinf) * 100

	// --- Header ---
	fmt.Println("\n\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    Materi 3                            ║")
	fmt.Println("║                Fungsi & Deret 2                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	// Output
	fmt.Println("Input Paket: a=9, r=0.1, N=14")
	fmt.Printf("\nSum Berhingga S(%d): %.2f\n", N, Sn)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", Sinf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", kedekatan)
}

// Soal 5: Deret suku (dipisah dari Soal 6)
func fd1() {
	// Parameter contoh
	c1 := 2
	c2 := 4
	n := 10
	//Suku awal//
	a0 := 1
	a1 := 1

	// --- Tampilkan Header ---
	fmt.Println("\n\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    Materi 3                            ║")
	fmt.Println("║                Fungsi & Deret 1                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	fmt.Printf("INPUT: C1=%d, C2=%d, N=%d\n", c1, c2, n)
	fmt.Println("Proses Perhitungan:")
	fmt.Printf("Suku 0:  %d | Suku 1:  %d", a0, a1)

	// Suku ke 0 dan 1
	if n == 0 {
		fmt.Printf("\nHASIL AKHIR Suku ke-0: %d\n", a0)
		return
	}
	if n == 1 {
		fmt.Printf("\nHASIL AKHIR Suku ke-1: %d\n", a1)
		return
	}

	//Operasi perhitungan suku
	var an int
	for i := 2; i <= n; i++ {

		an = (c1 * a1) + (c2 * a0)

		fmt.Printf(" | Suku %d:  %d", i, an)

		// Mengganti nilai suku sebelumnya
		a0 = a1
		a1 = an
	}

	// Hasil Akhir
	fmt.Printf("\nHASIL AKHIR Suku ke-%d:  %d\n", n, an)
}
