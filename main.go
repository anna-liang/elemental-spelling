package main

import (
	elem "elemental-spelling/element"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	var spelling, result []elem.Element
	var spelled string
	var spellReq bool
	elem.ImportElements()

	templates := template.Must(template.ParseFiles("templates/index.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spellReq = false
		result = nil
		spelling = nil
		if input := r.FormValue("input"); input != "" {
			spellReq = true
			spelling = elem.Spell(input, spelling[:])
			for _, element := range spelling {
				if strings.ToLower(spelled) != strings.ToLower(input) {
					spelled += element.Symbol
					result = append(result[:], element)
					input = input[len(element.Symbol):]
				}
			}
		}
		data := make(map[string]interface{})
		data["result"] = result
		data["spellReq"] = spellReq
		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Listening on localhost:8080...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
