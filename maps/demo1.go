package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["table"])
	fmt.Println("Eleman sayısı:", len(sozluk))
	fmt.Println("Sözlük:", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman sayısı:", len(sozluk))
	fmt.Println("Sözlük:", sozluk)

	deger, varMi := sozluk["bag"]
	fmt.Println(deger)
	fmt.Println("Sözlükte olma durumu:", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "plate": "tabak"}
	fmt.Println(sozluk2)
}
