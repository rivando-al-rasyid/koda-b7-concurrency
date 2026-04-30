package square

func Checkeven(in <-chan int, even chan<- int) {
	defer close(even)
	for n := range in {
		if n%2 == 0 {
			even <- n
		}
	}
}
