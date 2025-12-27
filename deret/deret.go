package deret

// Deret menghitung relasi rekurens iteratif
// a(n) = c1 * a(n-1) + c2 * a(n-2)
// dengan a(0) = 1, a(1) = 1
func Deret(c1, c2, n int) ([]int, int) {
	if n < 0 {
		return nil, 0
	}
	// Jika n = 0 cukup 1 elemen
	if n == 0 {
		return []int{1}, 1
	}

	// Buat slice untuk menyimpan semua suku
	a := make([]int, n+1)

	// Minimal butuh a0 dan a1
	a[0] = 1
	a[1] = 1

	for i := 2; i <= n; i++ {
		a[i] = c1*a[i-1] + c2*a[i-2]
	}
	return a, a[n]
}
