package main

import (
	"fmt"

	// "github.com/tenntenn/greeting"
	"github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do)
	fmt.Println(greeting.Do(time.Now()))
}
