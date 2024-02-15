package pkg

import (
	"fmt"
	"html/template"
	"image"
	"net/http"
	"strconv"

	"github.com/LQuatre/hangman-classic/core"
)

type TemplateData struct {
	Word        string
	Attempts    int
	WordLength  int
	LetterUsed  []string
	Info 	  	string
	Image 		string
}

var HangmanData core.HangManData
var usedLetters []string
var displayInfo string
var displayWord string
var myImage image.Image

func Run() {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	
	fs = http.FileServer(http.Dir("web/assets/img"))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", fs))


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

	for i := 0; i <= 10; i++ {
		http.HandleFunc(fmt.Sprintf("/assets/img/hangman%d.png", i), func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, fmt.Sprintf("web/assets/img/hangman%d.png", i))
		})
	}

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		tmpl, err := template.ParseFiles("web/template/play.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if method == "POST" {
			r.ParseForm()
			difficulty := r.Form.Get("difficulty")
			usedLetters = []string{}
			HangmanData = Init(difficulty)
		} else if method == "GET" {
			text := r.URL.Query().Get("text")
			displayInfo, *tmpl, usedLetters, HangmanData, err = ThisHangman(text, usedLetters, HangmanData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		displayWord = ""
		for i := 0; i < len(HangmanData.Word); i++ {
			displayWord += string(HangmanData.Word[i]) + " "
		}

		urlImage := "web/assets/img/hangman"+strconv.Itoa(0+(0-HangmanData.Attempts))+".png"
	

		data := TemplateData{
			Word: 	  displayWord,
			Attempts: HangmanData.Attempts,
			WordLength: HangmanData.WordLength,
			LetterUsed: usedLetters,
			Info: displayInfo,
			Image: urlImage,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
