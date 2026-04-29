package square

func Checkeven(in <-chan int, out chan<- int) {
	defer close(out)
	for n := range in {
		if n%2 == 0 {
			out <- n
		}
	}
}
