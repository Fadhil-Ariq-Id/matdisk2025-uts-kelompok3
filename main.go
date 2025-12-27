package main

import (
	"fmt"
	"math/rand"
	"time"

	"matdisk2025-uts-kelompok3/himpunan"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// ===== SOAL 1 =====
	N := 105
	limit := 20

	A := randomSlice(3, limit)
	B := randomSlice(3, limit)
	C := randomSlice(2, limit)

	fmt.Println("Generating Sets... (N=105)")
	fmt.Println("A:", A, "| B:", B, "| C:", C)

	R := himpunan.HitungR(A, B, C, N)
	fmt.Println("Hasil Operasi R:", R)

	filter := himpunan.FilterGenapKurangDari(R, N)
	fmt.Printf("Hasil Filter (Genap < %d): %v\n", N/4, filter)
	fmt.Println("Total Elemen:", len(filter))

	// ===== SOAL 2 =====
	fmt.Println("\nSoal 2: Analisis Pasangan Subset")

	M := 9
	K := 24
	S := randomSlice(M, 30)
	fmt.Println("Set S:", S, "| Target K:", K)

	pairs := himpunan.CariPasanganSubset(S, K)

	for i, p := range pairs {
		fmt.Printf("%d. {%d, %d} (Sum=%d)\n", i+1, p.X, p.Y, p.Sum)
	}
	fmt.Println("Total:", len(pairs), "Pasangan")
}

func randomSlice(n, limit int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(limit) + 1
	}
	return arr
}
