package interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	//dogrula(5)
	var sayi1 interface{} = 5 // true
	dogrula(sayi1)

	var sayi2 interface{} = "Sayi" // tip dönüşümü yapamadığı için -> false
	dogrula(sayi2)
}
