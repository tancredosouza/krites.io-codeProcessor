package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	enableCors(&w)
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		log.Println("Recebi uma request!")
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		log.Println(string(bodyBytes))
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}