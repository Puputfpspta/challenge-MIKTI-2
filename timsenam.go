package main

import (
	"fmt"
)
 //menghitung rata rata skor dari sebuah tim
func hitungRataRata(skors []int) float64 {
	total := 0
	for _, skor := range skors {
		total += skor
	}
	return float64(total) / float64(len(skors))
}

func main() {
	// Data Uji
	skorsLumbalumba := [][]int{
		{96, 108, 89},  // Data 1
		{97, 112, 101}, // Data Bonus 1
		{97, 112, 101}, // Data Bonus 2
	}
	skorsKoala := [][]int{
		{88, 91, 110}, // Data 1
		{109, 95, 123}, // Data Bonus 1
		{109, 95, 106}, // Data Bonus 2
	}

	// Perulangan untuk setiap data uji
	for i, skors := range skorsLumbalumba {
		// Hitung skor rata-rata untuk setiap tim
		rataRataLumbaLumba := hitungRataRata(skors)
		rataRataKoala := hitungRataRata(skorsKoala[i])

		fmt.Printf("Data %d:\n", i+1)
		fmt.Printf("Skor rata-rata Lumba-lumba: %.2f\n", rataRataLumbaLumba)
		fmt.Printf("Skor rata-rata Koala: %.2f\n", rataRataKoala)

		// Menentukan pemenang kompetisi
		if rataRataLumbaLumba > rataRataKoala {
			fmt.Println("Lumba-lumba menang.")
		} else if rataRataKoala > rataRataLumbaLumba {
			fmt.Println("Koala menang.")
		} else {
			// Handle hasil seri
			if rataRataLumbaLumba >= 100 && rataRataKoala >= 100 {
				fmt.Println("Hasil seri.")
			} else {
				fmt.Println("Tidak ada pemenang.")
			}
		}

		fmt.Println()
	}
}
