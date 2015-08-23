package palantir

import (
	"html/template"
	"io"
	"time"
)

func tmpl(w io.Writer, contentFile string, data interface{}) error {
	t := template.New("root.tpl")
	t.Funcs(template.FuncMap{
		"formatDate": func(data interface{}) (string, error) {
			epoch := data.(int64)
			loc, err := time.LoadLocation("Europe/Paris")
			if err != nil {
				return "", err
			}
			return time.Unix(epoch, 0).In(loc).Format("02/01/2006 15:04:05"), nil
		},
	},
	)
	t, err := t.ParseFiles("templates/root.tpl", "templates/"+contentFile)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}
