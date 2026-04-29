package square

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Generatenum(ch chan<- int) {
	defer close(ch)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("input angka: ")
	scanner.Scan()
	n := strings.TrimSpace(scanner.Text())

	num, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Error: input must be a number")
		return
	}

	for i := 1; i <= num; i++ {
		ch <- i
	}
}
