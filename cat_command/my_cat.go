package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var (
		option = flag.Bool("n", false, "-n option")
		j      = 1
		files  = os.Args
	)
	flag.Parse()
	fmt.Println(*option)
	for i := 1; i < (len(files)); i++ {
		//ファイルを探す
		err := filepath.Walk("./",
			func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() && info.Name() == files[i] {
					//ファイルのopenと読み込み&表示
					sf, err := os.Open(info.Name())
					if err != nil {
						return err
					}
					//openしたファイルの読み込み
					scanner := bufio.NewScanner(sf)
					for scanner.Scan() {
						if *option {
							fmt.Printf("%d: %s\n", j, scanner.Text())
						} else {
							fmt.Println(scanner.Text())
						}
						j++
					}
					if err = scanner.Err(); err != nil {
						fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
					}
					defer sf.Close()
				}
				return nil
			})
		if err != nil {
			fmt.Println(err)
		}
	}
}
