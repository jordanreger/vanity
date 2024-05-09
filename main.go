package vanity

import (
	"bytes"
	"html/template"
)

var body = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.Url}} {{.Type}} https://{{.GitUrl}}">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{.Url}}{{.Path}}">
<style>:root{color-scheme:light dark;}</style>
</head>
</html>`

type Vanity struct {
	Path   string
	Url    string
	GitUrl string
	Type   string
}

func Git(path string, root string, git_url string) string {
	st := Vanity{path, root, git_url, "git"}
	t, _ := template.New("import").Parse(body)
	var page bytes.Buffer

	err := t.ExecuteTemplate(&page, "import", st)
	if err != nil {
		panic(err)
	}
	return page.String()
}
