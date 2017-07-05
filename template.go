package c2d

import (
	"html/template"
	"os"
)

// CreateTemplate ...templateファイルを元にdnsmasq.templateを作成する
func (h *Host) CreateTemplate() error {
	f, err := os.Create(config.Output)
	if err != nil {
		return err
	}
	defer f.Close()

	t := template.Must(template.ParseFiles(config.TemplateFile))
	err = t.Execute(f, h)
	if err != nil {
		return err
	}
	return err
}
