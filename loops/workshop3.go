package loops

import "fmt"

func Workshop3() {
	//asal sayı mı
	sayi := 80
	tahmin := 0

	fmt.Println("Bir sayı giriniz...")
	fmt.Scanln(&tahmin)

	for sayi != tahmin {
		if tahmin < sayi {
			fmt.Println("Daha büyük bir sayı tahmin ediniz.")
			fmt.Scanln(&tahmin)
		} else if tahmin > sayi {
			fmt.Println("Daha küçük bir sayı tahmin ediniz.")
			fmt.Scanln(&tahmin)
		}
	}
	fmt.Println("Tebrikler!")
}
