package conditionals

import "fmt"

func Workshop1() {
	// üç adet sayıdan en büyüğünü ekrana yazdır
	var sayi1, sayi2, sayi3 int = 40, 30, 50
	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}

	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En büyük sayı: %v", enBuyuk)
}
