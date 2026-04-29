package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Kepada ")
	scanner.Scan()
	kepada := strings.TrimSpace(scanner.Text())

	fmt.Println("isi Pesan")
	scanner.Scan()
	pesan := strings.TrimSpace(scanner.Text())

	pesanfull := fmt.Sprintf("Kepada %s pesan %s", kepada, pesan)

	pesanSlice := []string{}
	pesanSlice = append(pesanSlice, pesanfull)

	pesanChan := make(chan []string)

	go papan.Pesan(pesanChan, pesanSlice)
	papan.CetakPesan(pesanChan)

}
