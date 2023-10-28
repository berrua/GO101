package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayı:", i)
		time.Sleep(time.Second * 1)
	}
}

func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayı:", i)
		time.Sleep(time.Second * 1)
	}
}
