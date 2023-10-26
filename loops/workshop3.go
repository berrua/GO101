package loops

import "fmt"

func Workshop3() {
	//asal sayı mı
	sayi := 0

	fmt.Println("Bir sayı giriniz...")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi == true {
		fmt.Println(sayi, "Asal sayıdır.")
	} else {
		fmt.Println(sayi, "Asal sayı değildir.")
	}

}
