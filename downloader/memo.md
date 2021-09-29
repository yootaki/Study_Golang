# 分割ダウンローダ
参考(https://github.com/Code-Hex/pget)

## 分割ダウンロードを行う
- Rangeアクセスを用いる
- いくつかのゴルーチンでダウンロードしてマージする
- エラー処理を工夫する
-- golang.org/x/sync/errgourpパッケージなどを使ってみる
- キャンセルが発生した場合の実装を行う

## 処理
[$ gget <url> <filename>] でダウンロードできるようにしたい
- 標準入力を受け取る
- パースしてurlを取得する
- filenameを取得＆fileopen
- httpリクエストを行う
- ファイルに書き込み（コピーを行う）

## 使用パッケージ等
### io
- func Copy(dst Writer, src Reader) (written int64, err error)
- func TeeReader(r Reader, w Writer) Reader

### os
- func OpenFile(name string, flag int, perm FileMode) (*File, error)

### http
- func (c *Client) Get(url string) (resp *Response, err error)
- func (srv *Server) Close() error
