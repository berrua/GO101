package defer_statements

import "fmt"

type product struct {
	name  string
	price int
}

func (p product) save() {
	fmt.Println("Kaydedildi:", p.name)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{name: "PC", price: 1000} // PC
	defer p.save()
	p = product{name: "Mouse", price: 100}
	//defer p.save()
	fmt.Println("Başarılı")
	fmt.Println(p.name)
	//defer p.save()
}
