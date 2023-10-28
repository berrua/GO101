package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"PC", 500, "XYZ"})
	fmt.Println(product{"PC", 500, "ABC"})
	fmt.Println(product{name: "PC", price: 500})
}

type product struct {
	name  string
	price float64
	brand string
}
