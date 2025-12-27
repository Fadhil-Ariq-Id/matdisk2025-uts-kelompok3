package himpunan

import "fmt"

// helper
func contains(arr []int, x int) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}

// A ⊕ B
func SelisihSimetri(A, B []int) []int {
	hasil := []int{}

	for _, a := range A {
		if !contains(B, a) {
			hasil = append(hasil, a)
		}
	}
	for _, b := range B {
		if !contains(A, b) {
			hasil = append(hasil, b)
		}
	}
	return hasil
}

// (A ⊕ B) \ C
func Selisih(A, C []int) []int {
	hasil := []int{}
	for _, a := range A {
		if !contains(C, a) {
			hasil = append(hasil, a)
		}
	}
	return hasil
}

// A ∩ C
func Irisan(A, C []int) []int {
	hasil := []int{}
	for _, a := range A {
		if contains(C, a) {
			hasil = append(hasil, a)
		}
	}
	return hasil
}

// Gabungan manual
func Gabungan(A, B []int) []int {
	hasil := []int{}

	for _, a := range A {
		if !contains(hasil, a) {
			hasil = append(hasil, a)
		}
	}
	for _, b := range B {
		if !contains(hasil, b) {
			hasil = append(hasil, b)
		}
	}
	return hasil
}

// Filter Genap & < N/4
func FilterGenapKurangDari(R []int, N int) []int {
	hasil := []int{}
	batas := N / 4

	for _, v := range R {
		if v%2 == 0 && v < batas {
			hasil = append(hasil, v)
		}
	}
	return hasil
}

// === FUNGSI UTAMA SESUAI SOAL ===
func HitungR(A, B, C []int, N int) []int {
	xor := SelisihSimetri(A, B)
	diff := Selisih(xor, C)
	inter := Irisan(A, C)
	R := Gabungan(diff, inter)
	return R
}

/// === NOmor 2 ====

type Pasangan struct {
	X   int
	Y   int
	Sum int
}

func CariPasanganSubset(S []int, K int) []Pasangan {
	hasil := []Pasangan{}
	used := make(map[string]bool)

	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {

			// cegah pasangan x == y
			// Elemen S digenerate acak, tidak dibatasi N
			if S[i] == S[j] {
				continue
			}

			sum := S[i] + S[j]
			if sum > K {

				// kunci pasangan unik (urutan tidak penting)
				a, b := S[i], S[j]
				if a > b {
					a, b = b, a
				}
				key := fmt.Sprintf("%d-%d", a, b)

				if !used[key] {
					used[key] = true
					hasil = append(hasil, Pasangan{
						X:   a,
						Y:   b,
						Sum: sum,
					})
				}
			}
		}
	}
	return hasil
}
