/*
HTTPサーバを作成する
リクエストが来たらおみくじの結果を返す
乱数の種は1回だけ初期化する
HTTPサーバを起動する前に初期化する
*/

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

var tmpl = template.Must(template.New("msg").Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です</body></html>"))

type Result struct {
	Name string
	Omikuji string
}

/* http://localhost:8080?p=Gopher */
func handler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Name: r.FormValue("p"),
		Omikuji: omikuji(),
	}
	tmpl.Execute(w, result)
}

func omikuji() string {
	n := rand.Intn(7)

	switch n + 1 {
	case 6:
		return "大吉"
	case 5:
		return "中吉"
	case 4:
		return "小吉"
	case 3:
		return "吉"
	case 2:
		return "凶"
	default:
		return "大凶"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	go func() {
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	}()

	//http request
	resp, err := http.Get("http://localhost:8080?p=Gopher")
	if err != nil {fmt.Println("error")}
	defer resp.Body.Close()
	var p Result
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&p); err != nil {
		fmt.Println("decode error")
	}
	fmt.Println(p)
}
