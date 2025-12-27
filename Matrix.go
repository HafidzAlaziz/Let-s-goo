package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inisialisasi seed random (wajib supaya tiap run beda)
	rand.Seed(time.Now().UnixNano())

	// PARAMETER PAKET 8 → N = 2
	N := 2

	// --- Generate Matrix A ---
	A := make([][]int, N)
	for i := range A {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			A[i][j] = rand.Intn(20) + 1 // 1..20
		}
	}

	// --- Generate Matrix B ---
	B := make([][]int, N)
	for i := range B {
		B[i] = make([]int, N)
		for j := 0; j < N; j++ {
			B[i][j] = rand.Intn(20) + 1
		}
	}

	fmt.Printf("Matrices generated (%dx%d)...\n\n", N, N)
	fmt.Println("Matrix A:", A)
	fmt.Println("Matrix B:", B)

	// --- Perkalian Matriks R = A x B ---
	R := make([][]int, N)
	for i := 0; i < N; i++ {
		R[i] = make([]int, N)
		for j := 0; j < N; j++ {
			sum := 0
			for k := 0; k < N; k++ {
				sum += A[i][k] * B[k][j]
			}
			R[i][j] = sum
		}
	}

	fmt.Println("\nHasil Matriks R:", R)

	// --- Hitung Trace ---
	trace := 0
	for i := 0; i < N; i++ {
		trace += R[i][i]
	}

	fmt.Printf("\nTrace: %d\n", trace)
}
