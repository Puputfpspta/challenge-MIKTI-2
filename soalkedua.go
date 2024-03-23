package main

import (
	"fmt"
)

func main() {
	// Data Uji
	// Data 1
	beratMark1 := 78.0 // kg
	tinggiMark1 := 1.69 // m
	beratJohn1 := 92.0 // kg
	tinggiJohn1 := 1.95 // m

	// Data 2
	beratMark2 := 95.0 // kg
	tinggiMark2 := 1.88 // m
	beratJohn2 := 85.0 // kg
	tinggiJohn2 := 1.76 // m

	// Hitung BMI (Body Mass Index)
	bmiMark1 := beratMark1 / (tinggiMark1 * tinggiMark1)
	bmiJohn1 := beratJohn1 / (tinggiJohn1 * tinggiJohn1)

	bmiMark2 := beratMark2 / (tinggiMark2 * tinggiMark2)
	bmiJohn2 := beratJohn2 / (tinggiJohn2 * tinggiJohn2)

	// Bandingkan BMI mereka
	markHigherBMI1 := bmiMark1 > bmiJohn1
	markHigherBMI2 := bmiMark2 > bmiJohn2

	// Cetak hasil perbandingan
	fmt.Println("Data 1:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark1)
	fmt.Printf("BMI John: %.2f\n", bmiJohn1)
	fmt.Printf("Mark memiliki BMI lebih tinggi dari John: %t\n\n", markHigherBMI1)

	fmt.Println("Data 2:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMark2)
	fmt.Printf("BMI John: %.2f\n", bmiJohn2)
	fmt.Printf("Mark memiliki BMI lebih tinggi dari John: %t\n", markHigherBMI2)
}