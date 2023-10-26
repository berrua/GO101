package loops

import "fmt"

func Workshop2() {
	//sayı tahmin oyunu
	//1-3 çok iyi, 4-6 iyi, >6 fena değil
	sayi := 80
	tahmin := 0
	hak := 0

	fmt.Println("Bir sayı tahmin ediniz...")
	fmt.Scanln(&tahmin)
	hak++

	for sayi != tahmin {
		if tahmin < sayi {
			fmt.Println("Daha büyük bir sayı tahmin ediniz.")
			fmt.Scanln(&tahmin)
			hak++
		} else if tahmin > sayi {
			fmt.Println("Daha küçük bir sayı tahmin ediniz.")
			fmt.Scanln(&tahmin)
			hak++
		}
	}

	durum := ""
	if hak > 0 && hak <= 3 {
		durum = "Çok iyi"
	} else if hak <= 6 {
		durum = "İyi"
	} else {
		durum = "Fena değil"
	}

	fmt.Printf("Tebrikler! %v tahmin: %v", hak, durum)
}
