package papan

import "fmt"

func CetakPesan(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}
