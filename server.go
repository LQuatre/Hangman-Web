package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Default to serving index.html
		filePath := "index.html"

		// If the path is explicitly /hangman.html, serve that file instead.
		if r.URL.Path == "/hangman.html" {
			filePath = "hangman.html"
		}
		if r.URL.Path == "/play.html" {
			filePath = "play.html"
		}

		http.ServeFile(w, r, filePath)
	})

	http.HandleFunc("/assets/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/style.css")
	})

	http.HandleFunc("/assets/hangmanstyle.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/hangmanstyle.css")
	})

	http.HandleFunc("/assets/playstyle.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/playstyle.css")
	})

	http.HandleFunc("/assets/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/script.js")
	})

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		if method == "POST" {
			r.ParseForm()
			difficulty := r.Form.Get("difficulty")
			fmt.Println("POST request successful")
			fmt.Println("Form data: ", r.Form)
			fmt.Println(difficulty)
			template.Must(template.ParseFiles("play.html")).Execute(w, nil)
		} else if method == "GET" {
			fmt.Println("GET request successful")
		}

	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
