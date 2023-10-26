package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(sehirler)

	kopya := make([]string, len(sehirler))
	fmt.Println("Kopya:", kopya)
	copy(kopya, sehirler)
	fmt.Println("Kopya:", kopya)

	sehirler = append(sehirler, "Adana")
	fmt.Println(sehirler)
	fmt.Println("Kopya:", kopya)

	fmt.Println(sehirler[1:3]) //1-3
	fmt.Println(sehirler[1:])  //1-
	fmt.Println(sehirler[:2])  //-2
}
