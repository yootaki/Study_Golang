/* 入力例
% go run search_history.go
book candy apple banana pine
book
candy
apple
banana
pine
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//get args from stdin
	args := bufio.NewScanner(os.Stdin)
	args.Scan()

	//split args with space
	array := strings.Split(args.Text(), " ")

	//make map
	maps := make(map[string]int)

	//count
	for _, key := range array {
		if val, ok := maps[key]; ok {
			maps[key] = val + 1
		} else {
			maps[key] = 1
		}
	}

	//print
	for key, val := range maps {
		fmt.Printf("%s %d\n", key, val)
	}
}
