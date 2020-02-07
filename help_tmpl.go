package clop

import (
	"io"
	"strings"
	"text/template"
)

func init() {
	funcMap = template.FuncMap{
		"addSpace": AddSpace,
	}
}

var funcMap map[string]interface{}

func AddSpace(max, cur int) string {
	return strings.Repeat(" ", max-cur)
}

type showOption struct {
	Opt   string
	Usage string
	Env   string
}

type Help struct {
	ProcessName string
	Version     string
	About       string
	Flags       []showOption
	Options     []showOption
	Args        []showOption
	MaxNameLen  int
}

func (h *Help) output(w io.Writer) error {
	tmpl := newTemplate()
	return tmpl.Execute(w, *h)
}

var usageDefaultTmpl = `{{if gt (len .Version) 0}}{{.Version}}{{end}}
{{if gt (len .About) 0}}{{.About}}{{end}}
{{if or (gt (len .Flags) 0) (gt (len .Options) 0) (gt (len .Args) 0)}}Usage:
    {{if gt (len .ProcessName) 0}}{{.ProcessName}}{{end}} {{if gt (len .Flags) 0}}[Flags] {{end}}{{if gt (len .Options) 0}}[Options] {{end}}{{range $_, $flag := .Args}}{{$flag.Opt}} {{end}}{{end}}
{{$maxNameLen :=.MaxNameLen}}
{{if gt (len .Flags) 0 }}Flags:
{{range $_, $flag:= .Flags}}    {{addSpace $maxNameLen (len $flag.Opt)|printf "%s%s" $flag.Opt}}    {{$flag.Usage}} {{if gt (len $flag.Env) 0 }}[env: {{$flag.Env}}]{{end}}
{{end}}{{end}}
{{if gt (len .Options) 0 }}Options:
{{range $_, $flag:= .Options}}    {{addSpace $maxNameLen (len $flag.Opt)|printf "%s%s" $flag.Opt}}    {{$flag.Usage}} {{if gt (len $flag.Env) 0 }}[env: {{$flag.Env}}]{{end}}
{{end}}{{end}}
{{if gt (len .Args) 0}}Args:
{{range $_, $flag:= .Args}}    {{addSpace $maxNameLen (len $flag.Opt)|printf "%s%s" $flag.Opt}}    {{$flag.Usage}} {{if gt (len $flag.Env) 0 }}[env: {{$flag.Env}}]{{end}}
{{end}}{{end}}
`

func newTemplate() *template.Template {
	tmpl := usageDefaultTmpl
	return template.Must(template.New("clop-default-usage").Funcs(funcMap).Parse(tmpl))
}
