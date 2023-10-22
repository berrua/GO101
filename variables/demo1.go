package variables

import "fmt"

func Demo1() {
	var metin string = "hello"
	fmt.Println(metin)
	fmt.Print(metin)
	fmt.Println(metin)

	var sayi int = 10
	fmt.Println(sayi)
	fmt.Println(100 + sayi)

	var sayi2 float32 = 0.5
	fmt.Println(sayi2)
	fmt.Println(100 + sayi2)

	sayi3 := 5
	fmt.Println(sayi3)
	fmt.Printf("veri tipi: %T\n", sayi3) // formatlı yazdırma

	var durum bool
	var metin1 string = "hello"
	var metin2 string = "helloo"
	durum = metin1 == metin2
	fmt.Println(durum)
	fmt.Println(!durum)
}
