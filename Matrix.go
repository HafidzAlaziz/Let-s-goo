package main

import (
	"fmt"
	"math/rand"
	"strings"
)

//═══════════════════════════════════════════════════════════════════════════════
//                           SOAL 3: PERKALIAN MATRIKS & TRACE
//═══════════════════════════════════════════════════════════════════════════════

func Soal3() {
	// PARAMETER PAKET 8 → N = 2
	N := 2

	// --- Generate Matrix A ---
	A := make([][]int, N)
	for i := range A {
		A[i] = make([]int, N)
		for j := range N {
			A[i][j] = rand.Intn(20) + 1 // 1..20
		}
	}

	// --- Generate Matrix B ---
	B := make([][]int, N)
	for i := range B {
		B[i] = make([]int, N)
		for j := range N {
			B[i][j] = rand.Intn(20) + 1
		}
	}

	// --- Print Header ---
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    SOAL NO. 3                          ║")
	fmt.Println("║              Perkalian Matriks & Trace                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Printf("\nParameter: N = %d (Matriks %dx%d)\n\n", N, N, N)

	// --- Print Matrix A ---
	fmt.Println("Matrix A:")
	for i := range N {
		fmt.Print("  [ ")
		for j := range N {
			fmt.Printf("%3d ", A[i][j])
		}
		fmt.Println("]")
	}

	// --- Print Matrix B ---
	fmt.Println("\nMatrix B:")
	for i := range N {
		fmt.Print("  [ ")
		for j := range N {
			fmt.Printf("%3d ", B[i][j])
		}
		fmt.Println("]")
	}

	// --- Perkalian Matriks R = A x B ---
	R := make([][]int, N)
	for i := range N {
		R[i] = make([]int, N)
		for j := range N {
			sum := 0
			for k := range N {
				sum += A[i][k] * B[k][j]
			}
			R[i][j] = sum
		}
	}

	// --- Print Hasil Perkalian ---
	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Println("Hasil Perkalian: R = A × B")
	fmt.Println(strings.Repeat("─", 56))
	for i := range N {
		fmt.Print("  [ ")
		for j := range N {
			fmt.Printf("%4d ", R[i][j])
		}
		fmt.Println("]")
	}

	// --- Hitung Trace ---
	trace := 0
	for i := range N {
		trace += R[i][i]
	}

	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Printf("Trace (Jumlah Diagonal Utama): %d\n", trace)
	fmt.Println(strings.Repeat("─", 56))
}

//═══════════════════════════════════════════════════════════════════════════════
//                      SOAL 4: TRANSFORMASI BARIS & MAKSIMUM
//═══════════════════════════════════════════════════════════════════════════════

func Soal4() {
	// PARAMETER PAKET 8 → N = 3
	N := 3

	// --- Generate Matrix M ---
	M := make([][]int, N)
	for i := range M {
		M[i] = make([]int, N)
		for j := range N {
			M[i][j] = rand.Intn(20) + 1 // 1..20
		}
	}

	// --- Print Header ---
	fmt.Println("\n\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    SOAL NO. 4                          ║")
	fmt.Println("║              Transformasi Baris & Maksimum             ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Printf("\nParameter: N = %d (Matriks %dx%d)\n\n", N, N, N)

	// --- Print Matrix M Awal ---
	fmt.Println("Matrix M (Generated):")
	for i := range N {
		fmt.Print("  [ ")
		for j := range N {
			fmt.Printf("%3d ", M[i][j])
		}
		fmt.Println("]")
	}

	// --- Tukar Baris 0 dengan Baris N-1 ---
	fmt.Printf("\nMenukar Baris 0 dan %d...\n", N-1)
	M[0], M[N-1] = M[N-1], M[0]

	// --- Print Matrix M Setelah Ditukar ---
	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Println("Matrix M Terkini (Setelah Tukar Baris):")
	fmt.Println(strings.Repeat("─", 56))
	for i := range N {
		fmt.Print("  [ ")
		for j := range N {
			fmt.Printf("%3d ", M[i][j])
		}
		fmt.Println("]")
	}

	// --- Cari Nilai Maksimum ---
	maxVal := M[0][0]
	maxRow := 0
	maxCol := 0

	for i := range N {
		for j := range N {
			if M[i][j] > maxVal {
				maxVal = M[i][j]
				maxRow = i
				maxCol = j
			}
		}
	}

	// --- Print Hasil ---
	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n", maxVal, maxRow, maxCol)
	fmt.Println(strings.Repeat("─", 56))
}
