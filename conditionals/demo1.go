package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var harcama float64 = 100

	if harcama > hesap {
		fmt.Println("Yetersiz bakiye!")
	}

	if harcama <= hesap {
		fmt.Println("İşlem başarılı!")
		hesap = hesap - harcama
	}

	fmt.Println("Bitti. Hesapta bakiyesi: " + fmt.Sprintf("%v", hesap)) // format dönüşümü
	fmt.Printf("Bitti. Hesapta bakiyesi: %v", hesap)                    // formatlı yazdırma
}
