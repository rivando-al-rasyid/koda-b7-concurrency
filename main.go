package main

import (
	"fmt"
	"sync"
	"time"
)

func Barista(order string) {
	fmt.Printf("Barista mulai membuat: %s\n", order)

	time.Sleep(1 * time.Second)

	fmt.Printf("Barista  SELESAI membuat: %s\n", order)
}

func main() {
	var wg sync.WaitGroup
	orders := []string{"Espresso", "Latte", "Cappuccino"}

	for _, order := range orders {
		wg.Go(func() { Barista(order) })
	}

	wg.Wait()
	fmt.Println("=== Semua pesanan telah dibuat. Toko Tutup! ===")
}
