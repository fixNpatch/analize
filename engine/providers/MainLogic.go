package providers

import (
	"encoding/json"
	"fmt"
	"strings"
)

const ALPHABET_RU = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
const ALPHABET_EN = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHABET_DE = "AÄBCDEFGHIJKLMNOÖPQRSẞTUÜVWXYZ"

type MainLogic struct{}

func (ml *MainLogic) CountCharInText(data *string) (result string) {
	var Data string
	Data = string(*data)

	alphabet := map[string]float64{}
	resultData := map[string]string{}
	charCounter := 0

	fmt.Println([]byte(Data))

	// set NULLs to resultData
	for _, r := range []rune(ALPHABET_RU) {
		char := string(r)
		alphabet[char] = 0
	}

	// counting each Character and write to resultData
	for _, r := range []rune(Data) { // add []rune for Cyrillic chars
		currentChar := strings.ToUpper(string(r))
		alphabet[currentChar] = alphabet[currentChar] + 1.0
		charCounter++
	}

	// redefine resultData by dividing onto overall number of chars
	for i := range alphabet {
		if alphabet[i] <= 0 {
			resultData[i] = "0"
			continue
		}
		fmt.Println(i, alphabet[i], "/", charCounter, " == ", float64(alphabet[i])/float64(charCounter))
		resultData[i] = fmt.Sprintf("%.5f", float64(alphabet[i])/float64(charCounter))

	}

	// format resultData to JSON
	jsonString, err := json.Marshal(resultData)
	if err != nil {
		fmt.Println("Something wrong with converting to JSON::", err)
		return
	}
	return string(jsonString)
}
