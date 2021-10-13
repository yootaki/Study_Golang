/* 入力例
% go run search_history.go
5
book
candy
apple
book
candy

candy
book
apple
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := bufio.NewScanner(os.Stdin)
	args.Scan()
	args_len, err := strconv.Atoi(args.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	slice := make([]string, args_len)
	for i := 0; i < args_len; i++ {
		args.Scan()
		if pos := same_slice(slice, args.Text()); pos == -1 {
			slice, slice[0] = append(slice[:1], slice[0:]...), args.Text()
		} else {
			slice = append(slice[:pos], slice[pos+1:]...)
			slice, slice[0] = append(slice[:1], slice[0:]...), args.Text()
		}
	}

	for i := 0; i < args_len; i++ {
		if slice[i] == "" {
			break
		}
		fmt.Println(slice[i])
	}
}

func same_slice(slice []string, arg string) int {
	for v := range slice {
		if slice[v] == arg {
			return v
		}
	}
	return -1
}
