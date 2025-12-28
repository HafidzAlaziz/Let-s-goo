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
	// PARAMETER PAKET 8 → ukuranMatriks = 2 (berarti matriks 2x2)
	ukuranMatriks := 2

	// --- Buat Matriks A dengan angka random ---
	matriksA := make([][]int, ukuranMatriks)
	for baris := range matriksA {
		matriksA[baris] = make([]int, ukuranMatriks)
		for kolom := range ukuranMatriks {
			matriksA[baris][kolom] = rand.Intn(20) + 1 // Angka random 1-20
		}
	}

	// --- Buat Matriks B dengan angka random ---
	matriksB := make([][]int, ukuranMatriks)
	for baris := range matriksB {
		matriksB[baris] = make([]int, ukuranMatriks)
		for kolom := range ukuranMatriks {
			matriksB[baris][kolom] = rand.Intn(20) + 1 // Angka random 1-20
		}
	}

	// --- Tampilkan Header ---
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    SOAL NO. 3                          ║")
	fmt.Println("║              Perkalian Matriks & Trace                 ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Printf("\nParameter: N = %d (Matriks %dx%d)\n\n", ukuranMatriks, ukuranMatriks, ukuranMatriks)

	// --- Tampilkan Matriks A ---
	fmt.Println("Matrix A:")
	for baris := range ukuranMatriks {
		fmt.Print("  [ ")
		for kolom := range ukuranMatriks {
			fmt.Printf("%3d ", matriksA[baris][kolom])
		}
		fmt.Println("]")
	}

	// --- Tampilkan Matriks B ---
	fmt.Println("\nMatrix B:")
	for baris := range ukuranMatriks {
		fmt.Print("  [ ")
		for kolom := range ukuranMatriks {
			fmt.Printf("%3d ", matriksB[baris][kolom])
		}
		fmt.Println("]")
	}

	// --- Kalikan Matriks A x B = Hasil ---
	// Rumus: Hasil[i][j] = A[i][0]*B[0][j] + A[i][1]*B[1][j] + ...
	matriksHasil := make([][]int, ukuranMatriks)
	for baris := range ukuranMatriks {
		matriksHasil[baris] = make([]int, ukuranMatriks)
		for kolom := range ukuranMatriks {
			jumlah := 0
			for indeks := range ukuranMatriks {
				// Kalikan baris A dengan kolom B, lalu jumlahkan
				jumlah += matriksA[baris][indeks] * matriksB[indeks][kolom]
			}
			matriksHasil[baris][kolom] = jumlah
		}
	}

	// --- Tampilkan Hasil Perkalian ---
	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Println("Hasil Perkalian: R = A × B")
	fmt.Println(strings.Repeat("─", 56))
	for baris := range ukuranMatriks {
		fmt.Print("  [ ")
		for kolom := range ukuranMatriks {
			fmt.Printf("%4d ", matriksHasil[baris][kolom])
		}
		fmt.Println("]")
	}

	// --- Hitung Trace (Jumlah diagonal utama) ---
	// Trace = Hasil[0][0] + Hasil[1][1] + Hasil[2][2] + ...
	nilaiTrace := 0
	for diagonal := range ukuranMatriks {
		nilaiTrace += matriksHasil[diagonal][diagonal]
	}

	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Printf("Trace (Jumlah Diagonal Utama): %d\n", nilaiTrace)
	fmt.Println(strings.Repeat("─", 56))
}

//═══════════════════════════════════════════════════════════════════════════════
//                      SOAL 4: TRANSFORMASI BARIS & MAKSIMUM
//═══════════════════════════════════════════════════════════════════════════════

func Soal4() {
	// PARAMETER PAKET 8 → ukuranMatriks = 3 (berarti matriks 3x3)
	ukuranMatriks := 3

	// --- Buat Matriks M dengan angka random ---
	matriksM := make([][]int, ukuranMatriks)
	for baris := range matriksM {
		matriksM[baris] = make([]int, ukuranMatriks)
		for kolom := range ukuranMatriks {
			matriksM[baris][kolom] = rand.Intn(20) + 1 // Angka random 1-20
		}
	}

	// --- Tampilkan Header ---
	fmt.Println("\n\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    SOAL NO. 4                          ║")
	fmt.Println("║              Transformasi Baris & Maksimum             ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Printf("\nParameter: N = %d (Matriks %dx%d)\n\n", ukuranMatriks, ukuranMatriks, ukuranMatriks)

	// --- Tampilkan Matriks M Sebelum Ditukar ---
	fmt.Println("Matrix M (Generated):")
	for baris := range ukuranMatriks {
		fmt.Print("  [ ")
		for kolom := range ukuranMatriks {
			fmt.Printf("%3d ", matriksM[baris][kolom])
		}
		fmt.Println("]")
	}

	// --- Tukar Baris Pertama (0) dengan Baris Terakhir (N-1) ---
	fmt.Printf("\nMenukar Baris 0 dan %d...\n", ukuranMatriks-1)
	// Swap: tukar isi baris 0 dengan baris terakhir
	matriksM[0], matriksM[ukuranMatriks-1] = matriksM[ukuranMatriks-1], matriksM[0]

	// --- Tampilkan Matriks M Setelah Ditukar ---
	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Println("Matrix M Terkini (Setelah Tukar Baris):")
	fmt.Println(strings.Repeat("─", 56))
	for baris := range ukuranMatriks {
		fmt.Print("  [ ")
		for kolom := range ukuranMatriks {
			fmt.Printf("%3d ", matriksM[baris][kolom])
		}
		fmt.Println("]")
	}

	// --- Cari Nilai Terbesar di Matriks ---
	nilaiTerbesar := matriksM[0][0] // Awalnya anggap angka pertama yang terbesar
	barisTerbesar := 0              // Posisi baris dari angka terbesar
	kolomTerbesar := 0              // Posisi kolom dari angka terbesar

	// Loop semua elemen matriks untuk cari yang paling besar
	for baris := range ukuranMatriks {
		for kolom := range ukuranMatriks {
			// Jika ketemu angka yang lebih besar, update
			if matriksM[baris][kolom] > nilaiTerbesar {
				nilaiTerbesar = matriksM[baris][kolom] // Simpan angka terbesar
				barisTerbesar = baris                  // Simpan posisi barisnya
				kolomTerbesar = kolom                  // Simpan posisi kolomnya
			}
		}
	}

	// --- Tampilkan Hasil ---
	fmt.Println("\n" + strings.Repeat("─", 56))
	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n", nilaiTerbesar, barisTerbesar, kolomTerbesar)
	fmt.Println(strings.Repeat("─", 56))
}
