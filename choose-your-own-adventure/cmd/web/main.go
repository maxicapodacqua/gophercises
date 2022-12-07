package main

import (
	"encoding/json"
	"fmt"
	"github.com/maxicapodacqua/choose-your-own-adventure/internal"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	tmpl := template.Must(template.ParseFiles("./templates/story.gohtml"))

	f, _ := os.Open("./gopher.json")

	var allStories map[string]internal.Page

	byteVal, _ := ioutil.ReadAll(f)
	json.Unmarshal(byteVal, &allStories)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		arc := strings.Trim(request.RequestURI, "/")

		err := tmpl.Execute(writer, allStories[arc])
		if err != nil {
			panic(err)
		}
	})

	fmt.Printf("%#v", http.DefaultServeMux)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
