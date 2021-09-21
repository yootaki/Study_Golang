/*
# 1コードポイント(rune)ずつ読み込むScannerを作れ
 - 初期化時にio.Readerを渡す
 - bufio.Scannerと似た感じに使えること
 - エラー処理を纏める
*/

/*
runeとは？ -> 文字列を1文字分ずつ扱うことができる型
bufio.Scannerとは？ ->
runeずつ読み込む方法
初期化時にio.Readerを渡すとは？
エラー処理を縮める
*/

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode/utf8"
)

/* io.Readerの実装 */
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }
type RuneScanner struct {
	r   io.Reader
	buf [16]byte
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	//{r: r}はフィールド名を指定して値を渡している
	//引数で受け取ったReaderをio.Readerに渡して、
	//RuneScannerのアドレスを返している
	return &RuneScanner{r: r}
}

//Scanはメソッド?
func (s *RuneScanner) Scan() (rune, error) {
	//Readerからbufに読み込む
	n, err := s.r.Read(s.buf[:])
	if err != nil {
		return 0, err
	}

	//s.bufがからならエラー
	//bufの最初のUTF-8エンコードを解凍して、ルーンとその幅をバイト数で返す
	r, size := utf8.DecodeRune(s.buf[:n])
	if r == utf8.RuneError {
		return 0, errors.New("RuneError")
	}

	//utf8.DecodeRuneで読み込んだ分（1文字分？）進めてRuneScannerのrに再度入れる
	//読み込んでDecodeした1文字ぶんをreturnする
	s.r = io.MultiReader(bytes.NewReader(s.buf[size:n]), s.r)
	return r, nil
}

func main() {
	scanner := NewRuneScanner(strings.NewReader("hello, world"))
	for {
		//1文字ぶんずつDecodeして読み込む
		r, err := scanner.Scan()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(r)
		//stringにキャストするとちゃんと文字として表示される
		fmt.Println(string(r))
	}
}
