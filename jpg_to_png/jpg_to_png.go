package main

/*
# 次の仕様を満たすコマンドを作って下さい
1. ディレクトリを指定する
 - os.Argsで引数を受け取る
 - そのディレクトリをfilepath.Walkで探す
 - 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
2. jpgファイルを探す
3. jpgをpngに変換
 - ディレクトリ以下は再帰的に処理する
4. 上のjpgを探して変換する処理は再帰で行う <--------------未実装
 - 変換前と変換後の画像形式を指定できる（オプション）

# 以下を満たすように開発してください
 - mainパッケージと分離する
 - 自作パッケージと標準パッケージと準標準パッケージのみ使う
 -- 準標準パッケージ：golang.org/x以下のパッケージ
 - ユーザ定義型を作ってみる
 - GoDocを生成してみる <------------------------------未実装
 - Go Modulesを使ってみる
*/

/*
使用方法
$ go run jpg_to_png.go <変換したい画像が入ったディレクトリ>

例
$ go run jpg_to_png.go img_dir
*/

import (
	"fmt"
	"os"
	"path/filepath"

	conv "jpg_to_png/conversion_img"
)

func main() {
	//ユーザー定義型
	img_info := struct {
		dir_name string
		// before_fmt string
		// after_fmt string
	}{
		dir_name: os.Args[1],
		// before_fmt: os.Args[2],
		// after_fmt: os.Args[3],
	}

	//カレントディレクトリ内に指定されたディレクトリがあるか探す
	err := filepath.Walk("./",
	func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == img_info.dir_name {
			//変換するパッケージを呼び出す
			conv.Convert_files(info)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
