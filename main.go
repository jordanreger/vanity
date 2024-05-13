package vanity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var body = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.Url}} {{.Type}} https://{{.GitUrl}}">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{.Url}}">
<style>:root{color-scheme:light dark;}</style>
</head>
</html>`

type vanity struct {
	Path   string
	Url    string
	GitUrl string
	Type   string
}

type config struct {
	Host string `json:"host"`
	URLs []url  `json:"urls"`
}
type url struct {
	Pkg string `json:"pkg"`
	URL string `json:"url"`
}

var cfg config

func init() {
	config_file, err := os.ReadFile("vanity.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(config_file, &cfg)
	if err != nil {
		panic(err)
	}
}

func Serve(mux *http.ServeMux) {
	for _, url := range cfg.URLs {
		mux.Handle("/"+url.Pkg+"/", Handler())
	}
}

func Handler() http.Handler {
	t, _ := template.New("import").Parse(body)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Content-Type", "text/html;charset=utf-8")
		url := r.URL
		path := url.Path

		root_pkg := strings.Split(path, "/")[1]

		for _, url := range cfg.URLs {
			if root_pkg == url.Pkg {
				st := vanity{strings.TrimPrefix(path, root_pkg), cfg.Host + strings.TrimSuffix(path, "/"), url.URL, "git"}
				var page bytes.Buffer

				err := t.ExecuteTemplate(&page, "import", st)
				if err != nil {
					panic(err)
				}
				fmt.Fprint(w, page.String())
			}
		}
	})
}
