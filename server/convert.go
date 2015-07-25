package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/danhardman/image-to-table/app"
	"github.com/danhardman/image-to-table/utils"
)

//Convert ...
func Convert(w http.ResponseWriter, r *http.Request) {
	file, fh, err := r.FormFile("file")
	utils.PanicOnError(err)

	table, err := app.ImageToTable(file, fh)
	if err != nil {
		fmt.Fprintf(w, "Unable to decode file. Error: %v", err)
	} else {
		fp, err := filepath.Abs("public/templates/table.html")
		utils.PanicOnError(err)

		t, err := template.ParseFiles(fp)
		utils.PanicOnError(err)

		t.Execute(w, table)
	}
}
