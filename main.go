package main

import (
	"github.com/rivando-al-rasyid/koda-b7-concurrency/internals/papan"
)

func main() {
	// var wg sync.WaitGroup
	// orders := []string{"Espresso", "Latte", "Cappuccino"}

	// for _, order := range orders {
	// 	wg.Go(func() { cafe.Barista(order) })
	// }
	// wg.Wait()
	// fmt.Println("=== Semua pesanan telah dibuat. Toko Tutup! ===")

	// activity := []string{"mandi", "buat kopi", "menyiapkan sarapan"}

	// activityChan := make(chan []string)
	// activityChan <- activity

	// for _, kegiatan := range activity {
	// 	wg.Go(func() { activities.Pagi(kegiatan) })
	// }

	pesanChan := make(chan string)
	go pap.Pesan(pesanChan)
	papan.CetakPesan(pesanChan)

}
