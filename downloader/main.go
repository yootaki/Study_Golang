package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Progress struct {
	total int64
	size  int64
}

func (p *Progress) Write(data []byte) (int, error) {
	n := len(data)
	p.size += int64(n)

	fmt.Fprintf(os.Stdout, "%d/%d\n", p.size, p.total)

	return n, nil
}

func main() {
	//get url & filename
	flag.Parse()
	args := flag.Args()

	//file open
	fp, err := os.OpenFile(args[1], os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer fp.Close()

	var p Progress

	//httpパッケージを使用してHTTPリクエストを行う
	resp, err := http.Get(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	//resp.Bodyの合計
	p.total = resp.ContentLength

	//ファイルに書き込みながら、p.Write内も実行する
	io.Copy(fp, io.TeeReader(resp.Body, &p))
}
/*
ダウンロードはできるようになったが、課題要件は一つも満たしていない。
*/
