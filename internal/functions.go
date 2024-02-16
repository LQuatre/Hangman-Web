package game

import (
	"fmt"
	"html/template"
	"image"
	"math/rand"
	"os"
	"strings"

	"github.com/LQuatre/hangman-classic/core"
)

func Init(difficulty string) core.HangManData { //initialiser les données du jeu du pendu en fonction du niveau de difficulté.
	MyHangManData := core.HangManData{} //Initialise une variable MyHangManData de type core.HangManData qui contiendra les données du jeu.
	var openFile []string
	switch difficulty { //strcuture de contrôle switch sélectionnant le fichier de mots en fonction de difficulté
	case "easy":
		openFile = core.ReadFile("pkg/assets/Dictionnary/words.txt")
	case "medium":
		openFile = core.ReadFile("pkg/assets/Dictionnary/words2.txt")
	case "hard":
		openFile = core.ReadFile("pkg/assets/Dictionnary/words3.txt")
	}
	lenFile := len(openFile)
	random := rand.Intn(lenFile)
	MyHangManData.ToFind = openFile[random] // mot aléatoire mit dans la variable MyHangManData

	lettersList := strings.Split(MyHangManData.ToFind, "")
	MyHangManData.WordLength = len(lettersList)
	numberOfLettersToReveal := MyHangManData.WordLength/2 - 1
	LettersToReveal := lettersList[:numberOfLettersToReveal]

	MyHangManData.Word = ""
	for i := 0; i < MyHangManData.WordLength; i++ {
		MyHangManData.Word += "_"
	}
	for i := 0; i < numberOfLettersToReveal; i++ {
		MyHangManData.Word = MyHangManData.Word[:i] + LettersToReveal[i] + MyHangManData.Word[i+1:]
	}
	fmt.Printf("Word to find: %s\n", MyHangManData.ToFind)
	MyHangManData.Attempts = 10
	return MyHangManData
}

func LoadImage(path string) (image.Image, error) { //fonction qui charge une image
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func ThisHangman(text string, usedLetters []string, HangmanData core.HangManData) (string, template.Template, []string, core.HangManData, error) { 
	var displayInfo string
	tmpl, err := template.ParseFiles("web/template/play.html")
	if err != nil {
		return "", *tmpl, []string{}, core.HangManData{}, err
	}

	if text != "" {
		letterFoundused, _ := core.ContainsArray(usedLetters, text)
		if letterFoundused {
			displayInfo = "Letter already used"
		} else {
			usedLetters = append(usedLetters, text)
			if len(text) == 1 {
				letterFound, _ := core.Contains(HangmanData.ToFind, text)
				if letterFound {
					HangmanData.Word = RevealedLetters(HangmanData.Word, HangmanData.ToFind, text)
				} else {
					HangmanData.Attempts--					
				}
			} else {
				if text != HangmanData.ToFind {
					HangmanData.Attempts = HangmanData.Attempts - 2
				} else {
					HangmanData.Word = HangmanData.ToFind
				}
			}
		}
	} else {
		displayInfo = "Please enter a letter"
	}

	if HangmanData.Word == HangmanData.ToFind {
		tmpl, err = template.ParseFiles("web/template/win.html")
		if err != nil {
			return "", *tmpl, []string{}, core.HangManData{}, err
		}
	}

	if HangmanData.Attempts == 0 {
		tmpl, err = template.ParseFiles("web/template/loose.html")
		if err != nil {
			return "", *tmpl, []string{}, core.HangManData{}, err
		}
	}

	return displayInfo, *tmpl, usedLetters, HangmanData, nil
}

func RevealedLetters(Word string, ToFind string, Letter string) string { //fonction qui remplace les '_' par les lettres trouvées
	WordList := strings.Split(Word, "")
	ToFindList := strings.Split(ToFind, "")
	for i := 0; i < len(ToFindList); i++ {
		if ToFindList[i] == Letter {
			WordList[i] = Letter
		}
	}
	return strings.Join(WordList, "")
}