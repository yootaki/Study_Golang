package main

import (
	"fmt"
)

type MyInt int
func (n *MyInt) Inc() { *n++ }

func main() {
	var n MyInt
	fmt.Println(n)
	(&n).Inc()
	fmt.Println(n)
}
