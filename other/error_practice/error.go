/*
# Stringerインタフェースに変換する関数を作ろう
 - 任意の値をStringer型に変換する関数
 -- type Stringer{String() string}
 - 引数にempty interfaceを取り、Stringerとエラーを返す
 -- func ToStringer(v interface{}) (Stringer, error)
 - 返すエラー型はerrorインタフェースを実装したユーザ定義型にする
 - 実際に呼び出してみてエラー処理をしてみよう
 -- エラーが発生した場合はエラーが発生した旨を表示する
*/

package main

import (
	"fmt"
	"os"
)

type Stringer interface {
	String() string
}
func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("CastError")
}

type MyError string
func (e MyError) Error() string {
	return string(e)
}

type S string
func (s S) Stringer() string {
	return string(s)
}

func main() {
	v := S("hello")
	if s, err := ToStringer(v); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}
}
