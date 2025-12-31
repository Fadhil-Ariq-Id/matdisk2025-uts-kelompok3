# 🔴 POTENSI DANGER (Bukan Error, tapi Catatan)

⚠️ 1. Matrix Tidak Dinamis

Saat ini:

n := len(a)

Artinya:

Program asumsi matriks persegi

Tidak cek:

ukuran A == ukuran B

kolom A == baris B

🔴 Kalau dosen minta N×N dinamis dari input → ini bisa jadi masalah

TAPI ❗
👉 Soal contoh pakai 2×2 hardcode, jadi AMAN UNTUK SOAL 3

---

🔥 TO-DO LIST SOAL 4 (WAJIB)
✅ TO-DO LIST RESMI

☑️ Generate / siapkan matriks M (N×N)
☑️ Tampilkan matriks awal
☑️ Tukar baris 0 ↔ N-1
☑️ Tampilkan matriks setelah ditukar
☑️ Cari nilai maksimum
☑️ Simpan posisi (row, col)
☑️ Print hasil akhir

4️⃣ 🔧 DESIGN FUNGSI (matrix.go)
Fungsi Tugas
SwapRows(m [][]int, r1, r2 int) Tukar dua baris
MaxValue(m [][]int) Cari max + posisi
