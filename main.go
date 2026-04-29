package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/rivando-al-rasyid/koda-b7-concurrency/internals/activities"
	"github.com/rivando-al-rasyid/koda-b7-concurrency/internals/cafe"
	"github.com/rivando-al-rasyid/koda-b7-concurrency/internals/papan"
	"github.com/rivando-al-rasyid/koda-b7-concurrency/internals/square"
)

func main() {
	var wg sync.WaitGroup

	orders := []string{"Espresso", "Latte", "Cappuccino"}
	for _, order := range orders {
		wg.Add(1)
		go func(o string) {
			defer wg.Done()
			cafe.Barista(o)
		}(order)
	}
	wg.Wait()
	fmt.Println("=== Semua pesanan telah dibuat. Toko Tutup! ===")

	activity := []string{"mandi", "buat kopi", "menyiapkan sarapan"}
	for _, kegiatan := range activity {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			activities.Pagi(k)
		}(kegiatan)
	}
	wg.Wait()
	fmt.Println("=== Semua aktivitas pagi selesai! ===")

	scanner := bufio.NewScanner(os.Stdin)
	pesanSlice := []string{}
	for {
		fmt.Print("Kepada (ketik 'quit()' untuk berhenti): ")
		scanner.Scan()
		kepada := strings.TrimSpace(scanner.Text())
		if strings.ToLower(kepada) == "quit()" {
			break
		}
		fmt.Print("Isi Pesan: ")
		scanner.Scan()
		pesan := strings.TrimSpace(scanner.Text())
		pesanfull := fmt.Sprintf("Kepada %s pesan: %s", kepada, pesan)
		pesanSlice = append(pesanSlice, pesanfull)
	}

	if len(pesanSlice) > 0 {
		pesanChan := make(chan string)
		go papan.Pesan(pesanChan, pesanSlice)
		papan.CetakPesan(pesanChan)
	} else {
		fmt.Println("Tidak ada pesan yang dikirim.")
	}

	ch1 := make(chan int)
	ch2 := make(chan int)
	go square.Generatenum(ch1)
	go square.Checkeven(ch1, ch2)
	square.Makesquare(ch2)
}
