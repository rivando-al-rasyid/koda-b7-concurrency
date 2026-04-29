package cafe

import (
	"fmt"
	"time"
)

func Barista(order string) {
	fmt.Printf("Barista mulai membuat: %s\n", order)

	time.Sleep(1 * time.Second)

	fmt.Printf("Barista  SELESAI membuat: %s\n", order)
}
