package providers

import (
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
	return `
<!doctype html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<link rel="stylesheet" href="http://cdn.webix.com/edge/webix.css" type="text/css"> 
		<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.2/css/all.css">
		<style>` + m.getAssets("engine/resourses/skin.css") + `</style>
    	<script src="http://cdn.webix.com/edge/webix.js" type="text/javascript"></script>
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
		var parsedData, stringJSON, jsString string

		switch string(paramPlusData[0]) {
		case "e":
			parsedData = strings.TrimPrefix(paramPlusData, "englishTable:")
			stringJSON = logicModel.CountCharInText(&parsedData)
			jsString = `
			let data = ` + stringJSON + `;

			$$('en_datatable_part_1').clearAll();
			$$('en_datatable_part_2').clearAll();
			$$('en_datatable_part_3').clearAll();

			$$('en_datatable_part_1').parse(data);
			$$('en_datatable_part_2').parse(data);
			$$('en_datatable_part_3').parse(data);

		`
			break
		case "r":
			parsedData = strings.TrimPrefix(paramPlusData, "russianTable:")
			stringJSON = logicModel.CountRuneInText(&parsedData)
			jsString = `
			let data = ` + stringJSON + `;

			$$('ru_datatable_part_1').clearAll();
			$$('ru_datatable_part_2').clearAll();
			$$('ru_datatable_part_3').clearAll();

			$$('ru_datatable_part_1').parse(data);
			$$('ru_datatable_part_2').parse(data);
			$$('ru_datatable_part_3').parse(data);

		`
			break
		default:
			fmt.Println("BAD PARAM CAUGHT::EXIT")
			break
		}

		// save latest result to Storage
		LocalStorage = stringJSON

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
