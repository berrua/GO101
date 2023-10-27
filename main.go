package main

import (
	"fmt"
	"goexamples/functions"
)

func main() {
	//fmt.Print()
	//variables.Demo1()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//conditionals.Workshop1()
	//loops.Demo1()
	//loops.Demo2()
	//loops.Workshop1()
	//loops.Workshop2()
	//loops.Workshop3()
	//loops.Workshop4()
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	//arrays.Demo4()
	//slices.Demo1()
	//slices.Demo2()
	//functions.Hello("Admin")
	//var sonuc = functions.Topla(3, 2)
	//fmt.Println(sonuc * 10)
	sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(4, 6)
	fmt.Println("Toplama: ", sonuc1)
	fmt.Println("Çıkarma: ", sonuc2)
	fmt.Println("Çarpma: ", sonuc3)
	fmt.Println("Bölme: ", sonuc4)

}
