package providers

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MainLogic struct{}

func (ml *MainLogic) CountCharInText(data *string) (result string) {
	var Data string
	Data = string(*data)

	alphabet := map[string]float64{}

	fmt.Println([]byte(Data))

	for _, r := range []rune(Data) { // add []rune for Cyrillic chars
		currentChar := strings.ToUpper(string(r))
		alphabet[currentChar] = alphabet[currentChar] + 1.0
	}

	for i := range alphabet {
		fmt.Println(i, alphabet[i], "/", len(Data), " == ", float64(alphabet[i])/float64(len(Data)))
		alphabet[i] = float64(alphabet[i]) / float64(len(Data))
	}
	fmt.Println(alphabet)

	jsonString, err := json.Marshal(alphabet)
	if err != nil {
		fmt.Println("Something wrong with converting to JSON::", err)
		return
	}
	return string(jsonString)
}
