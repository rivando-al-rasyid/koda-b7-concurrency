package square

func Generatenum(ch chan<- int) {
	defer close(ch)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}
