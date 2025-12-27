package main

import (
	"fmt"
	"math/rand"
	"time"

	"matdisk2025-uts-kelompok3/deret"
	"matdisk2025-uts-kelompok3/himpunan"
	"matdisk2025-uts-kelompok3/matrix"
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

	// ===== SOAL 3 =====
	fmt.Println("\nSoal 3: Perkalian dan Trace Matriks")
	fmt.Println("Matrices generated (2x2)...")

	// Masukan matrix A dan B
	a := [][]int{
		{1, 2},
		{3, 4},
	}
	b := [][]int{
		{5, 6},
		{7, 8},
	}
	fmt.Printf("Matrix A: %v\nMatrix B: %v\n", a, b)

	// Hitung R = A x B
	r := matrix.Multiply(a, b)
	fmt.Printf("Hasil Matrix R: %v\n", r)

	// Hitung trace R
	tr := matrix.Trace(r)
	fmt.Printf("Trace R: %d\n", tr)

	// ===== SOAL 4 =====
	fmt.Println("\nSoal 4: Transformasi Baris")

	// Matriks M
	m := [][]int{
		{1, 2},
		{10, 5},
	}

	fmt.Println("Matrix M (Generated):", m)

	// Tukar baris 0 dan N-1
	matrix.SwapRows(m, 0, len(m)-1)
	fmt.Println("Menukar baris 0 dan", len(m)-1)
	fmt.Println("Matrix M Terkini:", m)

	// Cari nilai maksimum
	max, row, col := matrix.MaxValue(m)
	fmt.Printf("Nilai Maksimum %d ditemukan di posisi (%d, %d)\n", max, row, col)

	// ===== SOAL 5 =====
	fmt.Println("\nSoal 5: Relasi Rekurens Iteratif")

	var c1, c2, n int
	fmt.Print("Masukkan C1: ")
	fmt.Scan(&c1)
	fmt.Print("Masukkan C2: ")
	fmt.Scan(&c2)
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	// Panggil fungsi Deret dari package deret
	hasilDeret, hasilAkhir := deret.Deret(c1, c2, n)

	fmt.Printf("\nINPUT: C1=%d, C2=%d, N=%d\n", c1, c2, n)
	fmt.Println("Proses Perhitungan:")

	// Tampilkan semua suku
	for i, v := range hasilDeret {
		if i == len(hasilDeret)-1 {
			fmt.Printf("Suku %d: %d", i, v)
		} else {
			fmt.Printf("Suku %d: %d | ", i, v)
		}
	}
	fmt.Println()

	fmt.Printf("HASIL AKHIR Suku ke-%d: %d\n", n, hasilAkhir)
}

func randomSlice(n, limit int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(limit) + 1
	}
	return arr
}
