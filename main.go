package main

import (
	"goexamples/interfaces"
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

	//sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(4, 6)
	//fmt.Println("Toplama: ", sonuc1)
	//fmt.Println("Çıkarma: ", sonuc2)
	//fmt.Println("Çarpma: ", sonuc3)
	//fmt.Println("Bölme: ", sonuc4)

	//fmt.Println(functions.ToplaVariadic(1, 4, 5, 7, 9))
	//fmt.Println(functions.ToplaVariadic(1, 4))
	//fmt.Println(functions.ToplaVariadic())

	//sayilar := []int{1, 2, 3, 4}
	//fmt.Println(functions.ToplaVariadic(sayilar...))

	//maps.Demo1()

	//for_range.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()

	//sayi := 20
	//pointers.Demo1(&sayi)
	//fmt.Println("Main:", sayi)

	//sayilar := []int{1, 2, 3, 4}
	//pointers.Demo2(sayilar)
	//fmt.Println("Main:", sayilar[0])

	//structs.Demo1()
	//structs.Demo2()

	//go goroutines.CiftSayilar()
	//go goroutines.TekSayilar()
	//time.Sleep(time.Second * 5)
	//fmt.Println("Main bitti.")

	//go goroutines.Demo2()
	//time.Sleep(time.Second * 5)
	//go islem()
	//...

	//ciftSayiCn := make(chan int)
	//tekSayiCn := make(chan int)
	//go channels.CiftSayilar(ciftSayiCn)
	//go channels.TekSayilar(tekSayiCn)

	//ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	//carpim := ciftSayiToplam * tekSayiToplam
	//fmt.Println("Çarpım:", carpim)

	//interfaces.Demo1()
	interfaces.Demo2()
}
