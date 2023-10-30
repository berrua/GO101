package error_handling

import (
	"errors"
	"fmt"
)

func tahminEt(tahmin int) (string, error) {
	sayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}

	if tahmin > sayi {
		return "Daha küçük bir sayı giriniz.", nil
	}

	if tahmin < sayi {
		return "Daha büyük bir sayı giriniz.", nil
	}

	return "Bildiniz.", nil // hata yok
}

func Demo2() {
	mesaj, hata := tahminEt(500)
	fmt.Println(mesaj, hata)

	//fmt.Println(tahminEt(50))
	//fmt.Println(tahminEt(101))
	//fmt.Println(tahminEt(-1))
}
