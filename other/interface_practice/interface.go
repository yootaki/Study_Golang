/*
# Stringerインタフェース
 - String() string メソッドを持つインタフェースを作る
 - そして3つ以上Stringerインタフェースを実装する型を作る
# インタフェースを受け取る関数
 - Stringerインタフェースを引数で受け取る関数を作る
 - 受け取った値を上記の3つの具象型によって分岐し、具象型の型名と値を表示する
*/

package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (i I) String() string {
	return "I"
}

type B bool

func (b B) String() string {
	return "B"
}

type S string

func (s S) String() string {
	return "S"
}

//3つの具象型で分岐し、型名と値を表示する関数
func print_name_and_value(st Stringer) {
	switch v := st.(type) {
	case I:
		fmt.Printf("int, %s\n", v)
	case B:
		fmt.Printf("bool, %s\n", v)
	default:
		fmt.Printf("string, %s\n", v)
	}
}

func main() {
	var test Stringer
	test = I(42)
	print_name_and_value(test)
	test = B(true)
	print_name_and_value(test)
	test = S("hello")
	print_name_and_value(test)
}
