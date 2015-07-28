package xiidec

import (
    "bytes"
    "fmt"
    "net/http"
    "html/template"
)
var (
    templates = template.Must(template.ParseFiles(
        "template/index.html",
    ))
)
func handler_index(w http.ResponseWriter, r *http.Request) {
    b := &bytes.Buffer{}
    dat := struct {
        Title string
        Body  string
    }{
        Title: "test page",
        Body:  "てすとっと?",
    }
    if err := templates.ExecuteTemplate(b, "index.html", dat); err != nil {
        fmt.Fprint(w, err.Error())
    }
    b.WriteTo(w)
}
