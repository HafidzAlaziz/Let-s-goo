package main

import (
	"fmt"
	"math"
)

// Soal 6: Analisis Kedekatan Deret Geometri
func Soal6() {
	// Parameter Paket No. 8
	a := 9.0
	r := 0.1
	N := 14

	// Hitung jumlah N suku berhingga
	Sn := a * (1 - math.Pow(r, float64(N))) / (1 - r)

	// Hitung jumlah tak hingga
	Sinf := a / (1 - r)

	// Hitung persentase kedekatan
	kedekatan := (Sn / Sinf) * 100

	// --- Tampilkan Header ---
	fmt.Println("\n\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    SOAL NO. 6                          ║")
	fmt.Println("║                       Deret                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")

	// Output sesuai format Golang Style
	fmt.Println("Input Paket: a=9, r=0.1, N=14")
	fmt.Printf("\nSum Berhingga S(%d): %.2f\n", N, Sn)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", Sinf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", kedekatan)
}
