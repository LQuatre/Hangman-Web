package pkg

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/LQuatre/hangman-classic/core"
)

func Init(difficulty string) core.HangManData {
	MyHangManData := core.HangManData{}
	var openFile []string
	switch difficulty {
	case "easy":
		openFile = core.ReadFile("pkg/assets/Dictionnary/words.txt")
	case "medium":
		openFile = core.ReadFile("pkg/assets/Dictionnary/words2.txt")
	case "hard":
		openFile = core.ReadFile("pkg/assets/Dictionnary/words3.txt")
	}
	lenFile := len(openFile)
	random := rand.Intn(lenFile)
	MyHangManData.ToFind = openFile[random]

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
	return MyHangManData
}
