package example1

import "fmt"

// squaring numbers - pipeline with 3 stages

func gen(nums ...int) <-chan int {
	// a channel that seeds
	out := make(chan int)

	go func() {
		for _, n := range nums {
			// put the number into just create channel
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

// func main() {
// 	// Set up the pipeline.
// 	c := gen(2, 3)
// 	out := sq(c)

// 	// Consume the output.
// 	fmt.Println(<-out) // 4
// 	fmt.Println(<-out) // 9
// }

func main() {
	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}
