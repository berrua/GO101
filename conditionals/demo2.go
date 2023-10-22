package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var harcama float64 = 100

	if harcama > hesap {
		fmt.Println("Yetersiz bakiye!")
	} else if harcama == hesap {
		fmt.Println("İşlem başarılı!")
		fmt.Println("Bakiyeniz bitti.")
	} else {
		fmt.Println("İşlem başarılı!")
	}
}
