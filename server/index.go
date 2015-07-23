package server

import (
	"net/http"
	"path/filepath"

	"github.com/danhardman/image-to-table/utils"
)

//GetIndex ...
func GetIndex(w http.ResponseWriter, r *http.Request) {
	fp, err := filepath.Abs("public/index.html")
	utils.PanicOnError(err)

	http.ServeFile(w, r, fp)
}
