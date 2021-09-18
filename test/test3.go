package main

import (
	"fmt"
)

func main() {
	n, m := 10, 20
	swap2 := func(n, m *int) {
		*n, *m = *m, *n
	}
	swap2(&n, &m)
	fmt.Println(n, m)
}
