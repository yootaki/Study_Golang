package main

import (
	"fmt"
)

func main() {
	f := func(n int) bool { return(n % 2 == 0) }
	for i := 0; i <=100; i++ {
		print(i)
		if f(i) {
			fmt.Println("-偶数")
		} else {
			fmt.Println("-奇数")
		}
	}
}
