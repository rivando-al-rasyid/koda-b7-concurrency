package activities

import (
	"fmt"
	"time"
)

func Pagi(kegiatan string) {
	fmt.Printf("mulai Melakukan: %s\n", kegiatan)

	time.Sleep(1 * time.Second)

	fmt.Printf("SELESAI Melakukan: %s\n", kegiatan)
}
