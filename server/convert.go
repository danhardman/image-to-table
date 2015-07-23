package server

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/danhardman/image-to-table/app"
	"github.com/danhardman/image-to-table/utils"
)

//Convert ...
func Convert(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")

	table := app.ImageToTable(file)
	fp, err := filepath.Abs("templates/table.html")
	utils.PanicOnError(err)

	t, err := template.ParseFiles(fp)
	utils.PanicOnError(err)

	t.Execute(w, table)
}
