package conversionimg

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

//変換したいファイルが入っているディレクトリの情報がdir_info
func Convert_files(dir_info os.FileInfo) {
	err := filepath.Walk(dir_info.Name(),
	func(path string, file_info os.FileInfo, err error) error {
		if !file_info.IsDir() && strings.Contains(file_info.Name(), ".jpg") {
			/* ファイルが見つかるので変換する */
			//ファイルオープン
			fmt.Println(path)
			fmt.Println(file_info.Name())
			file, _ := os.Open(path)
			defer file.Close()
			//ファイルオブジェクトを画像オブジェクトに変換
			img, _ := jpeg.Decode(file)
			//出力ファイルを生成
			out_path := strings.Replace(path, ".jpg", ".png", 1)
			out, _ := os.Create(out_path)
			defer out.Close()
			//pngに変換
			png.Encode(out, img)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
