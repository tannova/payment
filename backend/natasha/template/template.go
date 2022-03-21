package template

import (
	"io/ioutil"
	"log"
	"path"
	"text/template"
)

const _dir = "/app/templates"

var _template *template.Template

func init() {
	files, err := ioutil.ReadDir(_dir)
	if err != nil {
		log.Println("could not read template dir", err)
	}

	var filenames []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filenames = append(filenames, path.Join(_dir, f.Name()))
	}

	_template, err = template.ParseFiles(filenames...)
	if err != nil {
		log.Println("could not parse template files", err)
	}
}

func Lookup(name string) *template.Template {
	if _template == nil {
		return nil
	}

	return _template.Lookup(name)
}
