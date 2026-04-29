package papan

func Pesan(ch chan<- string, pesanSlice []string) {
	for _, pesan := range pesanSlice {
		ch <- pesan
	}
	close(ch)
}
