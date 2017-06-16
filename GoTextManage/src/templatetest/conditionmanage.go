package main

import (
	"os"
	"text/template"
)

type People struct {
	Username string
	Gender string
}

func (p People)String() string {
	return "username:" + p.Username + "  gender:" + p.Gender
}

func main() {

	info := map[string]bool{
		"Han Meimei": true,
		"LiLei": false,
	}
	t := template.Must(template.New("test").Parse(`Married: Han Meimei:{{index . "Han Meimei"}}; Li Lei:{{.LiLei}}`))
	t.Execute(os.Stdout, info)

	tEmpty := template.New("hello")
	tEmpty = template.Must(tEmpty.Parse("{{.}}"))
	tEmpty.Execute(os.Stdout, People{Username:"username", Gender:"men"})

	//tEmpty := template.New("template test")
	//tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	//tEmpty.Execute(os.Stdout, nil)
	//
	//tWithValue := template.New("template test")
	//tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	//tWithValue.Execute(os.Stdout, nil)
	//
	//tIfElse := template.New("template test")
	//tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	//tIfElse.Execute(os.Stdout, nil)
}
