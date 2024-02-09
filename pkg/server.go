package pkg

import (
	"fmt"
	"net/http"
	"text/template"
)

func Run() {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Url path: ", r.URL.Path)

		// Default to serving index.html
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, "web/template/index.html")
			fmt.Println("Index.html served")
		}

		// If the path is explicitly /hangman.html, serve that file instead.
		if r.URL.Path == "/hangman.html" || r.URL.Path == "/web/template/hangman.html" {
			http.ServeFile(w, r, "web/template/hangman.html")
			fmt.Println("Hangman.html served")
		}
		if r.URL.Path == "/play.html" || r.URL.Path == "/web/template/play.html" {
			http.ServeFile(w, r, "web/template/play.html")
			fmt.Println("Play.html served")
		}
	})

	http.HandleFunc("/assets/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/assets/style.css")
	})

	http.HandleFunc("/assets/hangmanstyle.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/assets/hangmanstyle.css")
	})

	http.HandleFunc("/assets/playstyle.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/assets/playstyle.css")
	})

	http.HandleFunc("/assets/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/assets/script.js")
	})

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		if method == "POST" {
			r.ParseForm()
			difficulty := r.Form.Get("difficulty")
			fmt.Println("POST request successful")
			fmt.Println("Form data: ", r.Form)
			fmt.Println(difficulty)

			template.Must(template.ParseFiles("web/template/play.html")).Execute(w, nil)
		} else if method == "GET" {
			fmt.Println("GET request successful")
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
