package conditionals

import "fmt"

func Workshop1() {
	// üç adet sayıdan en büyüğünü ekrana yazdır
	var sayi1, sayi2, sayi3 int = 40, 30, 50

	if sayi1 > sayi2 && sayi1 > sayi3 {
		fmt.Printf("En büyük sayı: %v", sayi1)
	} else if sayi2 > sayi1 && sayi2 > sayi3 {
		fmt.Printf("En büyük sayı: %v", sayi2)
	} else {
		fmt.Printf("En büyük sayı: %v", sayi3)
	}
}
