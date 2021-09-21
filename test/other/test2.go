package main

import (
	"fmt"
)

func main() {
	var n,m int
	swap := func(n,m int) (n2,m2 int) {
		n2, m2 = m, n
		return
	}
	n, m = swap(10, 20)
	fmt.Println(n, m)
}
