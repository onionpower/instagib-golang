package boringpipelines

func Gen(n ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}
