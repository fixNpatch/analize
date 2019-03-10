package providers

import (
	"encoding/json"
	"fmt"
	"github.com/zserge/webview"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var LocalStorage string

type WindowModel struct{}

func NewWindowModel() *WindowModel {
	return &WindowModel{}
}

func (m *WindowModel) getAssets(src string) string {
	// read file
	dat, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("Error while reading file", err)
	}
	return string(dat)
}

func (m *WindowModel) IndexHTML() string {
	//	return `
	//<!doctype html>
	//<html lang="en">
	//	<head>
	//		<meta charset="UTF-8">
	//		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	//		<link rel="stylesheet" href="http://cdn.webix.com/edge/webix.css" type="text/css">
	//		<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.2/css/all.css">
	//		<style>` + m.getAssets("engine/resourses/skin.css") + `</style>
	//    	<script src="http://cdn.webix.com/edge/webix.js" type="text/javascript"></script>
	//		<script>` + m.getAssets("engine/resourses/index.js") + `</script>
	//	</head>
	//	<body></body>
	//</html>
	//`

	// OFFLINE USAGE

	return `
		<!doctype html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.2/css/all.css">
			<style> ` + m.getAssets("engine/resourses/webix.css") + `</style>
			<style>` + m.getAssets("engine/resourses/skin.css") + `</style>
    	    <script> ` + m.getAssets("engine/resourses/webix.js") + `</script>
			<script>` + m.getAssets("engine/resourses/index.js") + `</script>
		</head>
		<body></body>
	</html>
`

}

func (m *WindowModel) HandleRPC(w *webview.WebView, data *string) {
	wb := *w
	dt := *data
	switch {
	// Close application
	case dt == "close":
		wb.Terminate()

	// Get changed value of text
	case strings.HasPrefix(dt, "push_table:"):
		logicModel := new(MainLogic)

		fmt.Println(dt)

		paramPlusData := strings.TrimPrefix(dt, "push_table:")
		var parsedData, jsString string

		switch string(paramPlusData[0]) {
		case "e":
			parsedData = strings.TrimPrefix(paramPlusData, "englishTable:")
			stringJSON, headerOrder, rawData := logicModel.CountRuneInText(&parsedData, "english")
			fmt.Println("ResultData::", string(stringJSON))
			var firstPart, secondPart, thirdPart []interface{}

			for i := 0; i < len(headerOrder); i++ {
				if i < 22 {
					if i > 10 {
						secondPart = append(secondPart, headerOrder[i])
					} else {
						firstPart = append(firstPart, headerOrder[i])
					}
				} else {
					thirdPart = append(thirdPart, headerOrder[i])
				}
			}

			firstPartMarshaled, err := json.Marshal(firstPart)
			if err != nil {
				fmt.Println("Something wrong with converting to JSON::Part 1::", err)
				return
			}

			secondPartMarshaled, err := json.Marshal(secondPart)
			if err != nil {
				fmt.Println("Something wrong with converting to JSON::Part 2::", err)
				return
			}

			thirdPartMarshaled, err := json.Marshal(thirdPart)
			if err != nil {
				fmt.Println("Something wrong with converting to JSON::Part 3::", err)
				return
			}

			fmt.Println(string(firstPartMarshaled))
			fmt.Println(string(secondPartMarshaled))
			fmt.Println(string(thirdPartMarshaled))

			LocalStorage, err = m.saveToStorage(headerOrder, rawData)

			jsString = `
			let firstPart = ` + string(firstPartMarshaled) + `;
			let secondPart = ` + string(secondPartMarshaled) + `;
			let thirdPart = ` + string(thirdPartMarshaled) + `;
		

			let depart_1 = $$('en_datatable_part_1'),
				depart_2 = $$('en_datatable_part_2'),
				depart_3 = $$('en_datatable_part_3');

			let columns_header_part_1e = [];
			for(let i in firstPart) {
				columns_header_part_1e.push({
					id: firstPart[i],
					header: firstPart[i],
					width:70
				})
				
			}

			let columns_header_part_2e = [];
			for(let i in secondPart) {
				columns_header_part_2e.push({
					id: secondPart[i],
					header: secondPart[i],
					width:70
				})
			}

			let columns_header_part_3e = [];
			for(let i in thirdPart) {
				columns_header_part_3e.push({
					id: thirdPart[i],
					header: thirdPart[i],
					width:70
				})
			}

			depart_1.config.columns = [];
			depart_1.refreshColumns();

			depart_2.config.columns = [];
			depart_2.refreshColumns();

			depart_3.config.columns = [];
			depart_3.refreshColumns();


			depart_1.config.columns = columns_header_part_1e;
			depart_1.refreshColumns();

			depart_2.config.columns = columns_header_part_2e;
			depart_2.refreshColumns();

			depart_3.config.columns = columns_header_part_3e;
			depart_3.refreshColumns();



			let data = ` + stringJSON + `;

			depart_1.clearAll();
			depart_2.clearAll();
			depart_3.clearAll();

			depart_1.parse(data);
			depart_2.parse(data);
			depart_3.parse(data);
		`
			break
		case "r":
			parsedData = strings.TrimPrefix(paramPlusData, "russianTable:")
			stringJSON, headerOrder, rawData := logicModel.CountRuneInText(&parsedData, "russian")
			fmt.Println("ResultData::", string(stringJSON))
			var firstPart, secondPart, thirdPart []interface{}

			for i := 0; i < len(headerOrder); i++ {
				if i < 22 {
					if i > 10 {
						secondPart = append(secondPart, headerOrder[i])
					} else {
						firstPart = append(firstPart, headerOrder[i])
					}
				} else {
					thirdPart = append(thirdPart, headerOrder[i])
				}
			}

			firstPartMarshaled, err := json.Marshal(firstPart)
			if err != nil {
				fmt.Println("Something wrong with converting to JSON::Part 1::", err)
				return
			}

			secondPartMarshaled, err := json.Marshal(secondPart)
			if err != nil {
				fmt.Println("Something wrong with converting to JSON::Part 2::", err)
				return
			}

			thirdPartMarshaled, err := json.Marshal(thirdPart)
			if err != nil {
				fmt.Println("Something wrong with converting to JSON::Part 3::", err)
				return
			}

			fmt.Println(string(firstPartMarshaled))
			fmt.Println(string(secondPartMarshaled))
			fmt.Println(string(thirdPartMarshaled))

			LocalStorage, err = m.saveToStorage(headerOrder, rawData)

			jsString = `
			let firstPart = ` + string(firstPartMarshaled) + `;
			let secondPart = ` + string(secondPartMarshaled) + `;
			let thirdPart = ` + string(thirdPartMarshaled) + `;
		

			let dpart_1 = $$('ru_datatable_part_1'),
				dpart_2 = $$('ru_datatable_part_2'),
				dpart_3 = $$('ru_datatable_part_3');

			let columns_header_part_1 = [];
			for(let i in firstPart) {
				columns_header_part_1.push({
					id: firstPart[i],
					header: firstPart[i],
					width:70
				})
				
			}

			let columns_header_part_2 = [];
			for(let i in secondPart) {
				columns_header_part_2.push({
					id: secondPart[i],
					header: secondPart[i],
					width:70
				})
			}

			let columns_header_part_3 = [];
			for(let i in thirdPart) {
				columns_header_part_3.push({
					id: thirdPart[i],
					header: thirdPart[i],
					width:70
				})
			}

			dpart_1.config.columns = [];
			dpart_1.refreshColumns();

			dpart_2.config.columns = [];
			dpart_2.refreshColumns();

			dpart_3.config.columns = [];
			dpart_3.refreshColumns();


			dpart_1.config.columns = columns_header_part_1;
			dpart_1.refreshColumns();

			dpart_2.config.columns = columns_header_part_2;
			dpart_2.refreshColumns();

			dpart_3.config.columns = columns_header_part_3;
			dpart_3.refreshColumns();



			let data = ` + stringJSON + `;

			dpart_1.clearAll();
			dpart_2.clearAll();
			dpart_3.clearAll();

			dpart_1.parse(data);
			dpart_2.parse(data);
			dpart_3.parse(data);

		`
			break
		default:
			fmt.Println("BAD PARAM CAUGHT::EXIT")
			break
		}

		// save latest result to Storage
		fmt.Println("locstor::", LocalStorage)

		err := wb.Eval(jsString)
		if err != nil {
			fmt.Println("Error while executing JS::Push_table::", err)
			return
		}

	// Open file
	case dt == "open":
		log.Println("open") // log stamp
		// open Dialog window
		pathFile := wb.Dialog(webview.DialogTypeOpen, 0, "Open file", "") // absolute path to the file
		fmt.Println(pathFile)                                             // print resultedPath

		b, err := ioutil.ReadFile(pathFile) // just pass the file name
		if err != nil {
			fmt.Println("Catch error::Open file::Read", err)
			return
		}

		// Form JS
		jsString := `$$('resulted_text').setValue(` + strconv.Quote(string(b)) + `);`
		err = wb.Eval(jsString)
		if err != nil {
			log.Println("Catch error::Open file::OpenData", err)
			return
		}

	case dt == "save":
		log.Println("save") // log stamp
		// open Dialog window
		pathFile := wb.Dialog(webview.DialogTypeSave, webview.DialogFlagFile, "Save file", "") // absolute path to the file
		fmt.Println(pathFile)
		err := m.saveFile(pathFile)
		if err != nil {
			log.Println("Catch error::Writing file::", err)
			return
		}
	}
}

func (m *WindowModel) saveFile(pathFile string) (err error) {

	data := []byte(LocalStorage)

	// write the whole body at once
	err = ioutil.WriteFile(pathFile, data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("All good, there're no pointer's error")
	return nil
}

func (m *WindowModel) saveToStorage(order []interface{}, data map[string]string) (result string, err error) {
	for i := range order {
		value := "0"
		for j := range data {
			if order[i].(string) == j {
				value = data[j]
			}
		}
		result += order[i].(string) + ":" + value + "   "
	}

	return
}
