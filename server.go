package main

import (
	"fmt"
	"net/http"
)

func main() {
    fs := http.FileServer(http.Dir("assets"))
    http.Handle("/assets/", http.StripPrefix("/assets/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    // Serve CSS file
    http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "assets/style.css")
    })

    // Serve JS file
    http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "assets/script.js")
    })

    // Import CSS and JS files
    http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })
    http.HandleFunc("/assets/style.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "assets/style.css")
    })
    http.HandleFunc("/assets/script.js", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "assets/script.js")
    })

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
