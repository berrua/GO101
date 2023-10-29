package defer_statements

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		fmt.Println("Ç")
		return "Çift sayı."
	}

	if sayi%2 != 0 {
		fmt.Println("T")
		return "Tek sayı."
	}

	return "Belli değil."
}

func Test() {
	sonuc := Demo2(8)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti.")
}
