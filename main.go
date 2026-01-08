package main

import (
	"fmt"
	"strings"
)

//═══════════════════════════════════════════════════════════════════════════════
//                                  MAIN PROGRAM
//═══════════════════════════════════════════════════════════════════════════════

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║                    Kelompok H & L                      ║")
	fmt.Println("║________________________________________________________║")
	fmt.Println("║  Anggota :                                             ║")
	fmt.Println("║           1. Lutfian Falah          (251552010027)     ║")
	fmt.Println("║           2. Muhammad Hafidz Alaziz (251552010034)     ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	// himpunan
	h1()
	h2()

	// martrix

	Matrix1()
	Matrix2()

	// deret & fungsi
	fd1()
	fd2()

	// Footer
	fmt.Println("\n" + strings.Repeat("═", 56))
	fmt.Println("                 Matrix - Done Lah..")
	fmt.Println(strings.Repeat("═", 56))
}
