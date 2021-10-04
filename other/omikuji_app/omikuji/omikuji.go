package omikuji

import (
	"math/rand"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.New("msg").Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です</body></html>"))

type Result struct {
	Name    string
	Omikuji string
}

/* http://localhost:8080?p=Gopher */
func handler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Name:    r.FormValue("p"),
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
