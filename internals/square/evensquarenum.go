package square

func Makesquare(in <-chan int) {
	for n := range in {
		println(n * n)
	}
}
