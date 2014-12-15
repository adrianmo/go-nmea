package pid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	indexFilename = "index.html"
)

// An indexHandler handles file requests.
type indexHandler struct {
}

// ServeHTTP returns the contents of the requested file.
func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := path.Base(r.URL.Path)
	if filename == "/" {
		filename = indexFilename
	}
	fmt.Printf("Open: %s\n", filename)
	indexData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(w, "error reading %s: %s", indexFilename, err)
		return
	}
	fmt.Fprintf(w, "%s", indexData)
}
