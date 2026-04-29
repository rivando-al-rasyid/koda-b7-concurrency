package papan

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pesan(ch chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("isi Pesan")

	scanner.Scan()
	pesan := strings.TrimSpace(scanner.Text())

	ch <- fmt.Sprintf("pesan %s", pesan)

	close(ch)
}
