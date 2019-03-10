package providers

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const ALPHABET_RU = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
const ALPHABET_EN = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type MainLogic struct{}

type kvSupport struct { // structure to sort arrray
	Key   string
	Value string
}

type OrderedMap struct {
	Order []string
	Map   map[string]string
}

func (ml *MainLogic) CountRuneInText(data *string, param string) (result string, headerOrder []interface{}) {
	var alphabetConst string
	switch param {
	case "english":
		{
			alphabetConst = ALPHABET_EN
			break
		}
	case "russian":
		{
			alphabetConst = ALPHABET_RU
			break
		}
	default:
		fmt.Println("ERROR. Didn't get parameter. ---- CountRuneInText")
		return
	}

	fmt.Println("Begin func - CountRuneInText...")
	var Data string
	Data = string(*data)

	alphabet := map[string]float64{}
	resultData := map[string]string{}
	charCounter := 0

	fmt.Println("Begin setting nulls to resultData")
	// set NULLs to resultData
	for _, r := range []rune(alphabetConst) {
		char := string(r)
		alphabet[char] = 0
	}

	fmt.Println("Begin counting Char and write to resultData")
	// counting each Character and write to resultData
	for _, r := range []rune(Data) { // add []rune for Cyrillic chars
		currentChar := strings.ToUpper(string(r))
		if !checkChar(currentChar, alphabetConst) { // if Char doesn't belong to existing alphabet
			continue
		}
		alphabet[currentChar] = alphabet[currentChar] + 1.0
		charCounter++
	}

	fmt.Println("Redefine resultData by dividing onto overall number of chars")
	// redefine resultData by dividing onto overall number of chars
	for i := range alphabet {
		if alphabet[i] <= 0 {
			resultData[i] = "0"
			continue
		}
		resultData[i] = fmt.Sprintf("%.4f", float64(alphabet[i])/float64(charCounter))
	}

	order := ml.getOrder(resultData)
	fmt.Println("order::", order)

	// format resultData to JSON
	jsonString, err := json.Marshal(resultData)
	if err != nil {
		fmt.Println("Something wrong with converting to JSON::", err)
		return
	}

	return string(jsonString), order
}

func (ml MainLogic) getOrder(data map[string]string) (order []interface{}) {
	resultData := sortSlice(data)

	for _, item := range resultData {
		header := item.(kvSupport).Key
		order = append(order, header)
	}

	return
}

func sortSlice(data map[string]string) (resultData []interface{}) {
	var supportSlice []kvSupport
	for k, v := range data {
		supportSlice = append(supportSlice, kvSupport{k, v})
	}
	sort.SliceStable(supportSlice, func(i, j int) bool {
		return supportSlice[i].Key <= supportSlice[j].Key
	})

	// Bubble sort
	for i := 1; i < len(supportSlice); i++ {
		for j := 0; j < len(supportSlice)-i; j++ {
			if supportSlice[j].Value < supportSlice[j+1].Value {
				supportSlice[j], supportSlice[j+1] = supportSlice[j+1], supportSlice[j]
			}
		}
	}

	for _, kv := range supportSlice {
		resultData = append(resultData, kv)
	}
	return resultData
}

func checkChar(char string, alphabet string) bool {
	for _, r := range []rune(alphabet) {
		if char == string(r) {
			return true
		}
	}
	return false
}
