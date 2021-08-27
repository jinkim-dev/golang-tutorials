package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zerotr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zerotr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Print("pointer:", &i)
}
