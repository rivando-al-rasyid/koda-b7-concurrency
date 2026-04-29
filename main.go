package main

import (
	"fmt"
	"sync"

	"github.com/rivando-al-rasyid/koda-b7-concurrency/internals/activities"
)

func main() {
	var wg sync.WaitGroup
	// orders := []string{"Espresso", "Latte", "Cappuccino"}

	// for _, order := range orders {
	// 	wg.Go(func() { cafe.Barista(order) })
	// }

	// fmt.Println("=== Semua pesanan telah dibuat. Toko Tutup! ===")

	activity := []string{"mandi", "buat kopi", "menyiapkan sarapan"}
	for _, kegiatan := range activity {
		wg.Go(func() { activities.Pagi(kegiatan) })
	}

	wg.Wait()
	fmt.Println("=== Berangkat Kerja ===")

}
