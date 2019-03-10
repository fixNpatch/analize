package providers

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const ALPHABET_RU = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
const ALPHABET_EN = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ALPHABET_DE = "AÄBCDEFGHIJKLMNOÖPQRSẞTUÜVWXYZ"

type MainLogic struct{}

type kvSupport struct { // structure to sort arrray
	Key   string
	Value string
}

type OrderedMap struct {
	Order []string
	Map   map[string]string
}

func (ml *MainLogic) CountCharInText(data *string) (result string) {
	var Data string
	Data = string(*data)

	alphabet := map[string]float64{}
	resultData := map[string]string{}
	charCounter := 0

	// set NULLs to resultData
	for _, r := range []rune(ALPHABET_EN) {
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
		resultData[i] = fmt.Sprintf("%.4f", float64(alphabet[i])/float64(charCounter))
	}

	// format resultData to JSON
	jsonString, err := json.Marshal(resultData)
	if err != nil {
		fmt.Println("Something wrong with converting to JSON::", err)
		return
	}

	fmt.Println("CountCharInText::jsonString::finishFunc::", string(jsonString))
	return string(jsonString)
}

func (ml *MainLogic) CountRuneInText(data *string) (result string, headerOrder []interface{}) {
	var Data string
	Data = string(*data)

	alphabet := map[string]float64{}
	resultData := map[string]string{}
	charCounter := 0

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

	orderJson, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Something wrong with converting to JSON::", err)
		return
	}

	fmt.Println("CountRuneInText::orderJson::finishFunc::", string(orderJson))

	fmt.Println("CountRuneInText::jsonString::finishFunc::", string(jsonString))
	return string(jsonString), order
}

//func (ml *MainLogic) sortSlice(data map[string]string) map[string]string {
//
//	fmt.Println("Step 2 ::Counted Map::")
//	fmt.Println(data)
//	fmt.Println("______________________\n", time.Now().UTC(), "\n______________________")
//
//	var supportSlice []kvSupport
//	resultData := map[string]string{}
//	for k, v := range data {
//		supportSlice = append(supportSlice, kvSupport{k, v})
//	}
//
//	sort.Slice(supportSlice, func(i, j int) bool {
//		return supportSlice[i].Value > supportSlice[j].Value
//	})
//
//	for _, kv := range supportSlice {
//		fmt.Printf("%s, %s\n", kv.Key, kv.Value) // for testing. delete to production
//		resultData[string(kv.Key)] = string(kv.Value)
//	}
//	return resultData
//}

//// unknown sort
//func (ml *MainLogic) sortSlice(data map[string]string) (resultData []interface{}) {
//
//	fmt.Println("Step 2 ::Counted Map::")
//	fmt.Println(data)
//	fmt.Println("______________________\n", time.Now().UTC(), "\n______________________")
//
//	var supportSlice []kvSupport
//	for k, v := range data {
//		supportSlice = append(supportSlice, kvSupport{k, v})
//	}
//	sort.SliceStable(supportSlice, func(i, j int) bool {
//		return supportSlice[i].Key <= supportSlice[j].Key
//	})
//
//	// Bubble sort
//	for i := 1; i < len(supportSlice); i++ {
//		for j := 0; j < len(supportSlice)-i; j++ {
//			if supportSlice[j].Value < supportSlice[j+1].Value {
//				supportSlice[j], supportSlice[j+1] = supportSlice[j+1], supportSlice[j]
//			}
//		}
//	}
//
//	for _, kv := range supportSlice {
//		fmt.Printf("%s::%s\n", kv.Key, kv.Value) // for testing. delete to production
//		pair := string(kv.Key) + ":" + string(kv.Value)
//		resultData = append(resultData, pair)
//	}
//	return resultData
//}

// stable sort
//func (ml *MainLogic) sortSlice(data map[string]string) (resultData []interface{}) {
//
//	fmt.Println("Step 2 ::Counted Map::")
//	fmt.Println(data)
//	fmt.Println("______________________\n", time.Now().UTC(), "\n______________________")
//
//	var supportSlice []kvSupport
//	for k, v := range data {
//		supportSlice = append(supportSlice, kvSupport{k, v})
//	}
//	sort.SliceStable(supportSlice, func(i, j int) bool {
//		return supportSlice[i].Key <= supportSlice[j].Key
//	})
//
//	// Bubble sort
//	for i := 1; i < len(supportSlice); i++ {
//		for j := 0; j < len(supportSlice)-i; j++ {
//			if supportSlice[j].Value < supportSlice[j+1].Value {
//				supportSlice[j], supportSlice[j+1] = supportSlice[j+1], supportSlice[j]
//			}
//		}
//	}
//
//	for _, kv := range supportSlice {
//		fmt.Printf("%s::%s\n", kv.Key, kv.Value) // for testing. delete to production
//		resultData = append(resultData, kv)
//	}
//	return resultData
//}

func (ml MainLogic) getOrder(data map[string]string) (order []interface{}) {
	resultData := sortSlice(data)

	for _, item := range resultData {
		header := item.(kvSupport).Key
		fmt.Println(header)
		order = append(order, header)
	}

	return
}

// unknown2 sort
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
		fmt.Printf("%s::%s\n", kv.Key, kv.Value) // for testing. delete to production
		resultData = append(resultData, kv)
	}
	return resultData
}
