package papan

func Pesan(ch chan<- []string, pesanSlice []string) {
	ch <- pesanSlice

	close(ch)

}
