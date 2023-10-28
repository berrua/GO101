package goroutines

import (
	"fmt"
)

func Demo2() {
	for i := byte('a'); i <= byte('z'); i++ {
		fmt.Println(string(i))
	}
}
