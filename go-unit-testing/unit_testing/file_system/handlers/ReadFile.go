package handlers

import (
	"html/template"
	"net/http"
	"path"
)

func ReadFile(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles(path.Join("./templates/test", "readfile.html"))
	if err != nil {
		w.Write([]byte("Bad Request!" + err.Error()))
		return
	}

	data := struct {
		Success  string
		FilePath string
	}{
		Success:  "If you see this, means your template is working fine!",
		FilePath: path.Join("./templates/test", "readfile.html"),
	}

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		w.Write([]byte("Error executing go template" + err.Error()))
		return
	}
}
