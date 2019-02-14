package providers

import (
	"fmt"
	"github.com/zserge/webview"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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
		parsedData := (strings.TrimPrefix(dt, "push_table:"))
		logicModel := new(MainLogic)
		stringJSON := logicModel.CountCharInText(&parsedData)
		fmt.Println(stringJSON)

		jsString := `
			// let data = [{"Б":2}];
			// let dt = $$('datatable_id');
			// dt.parse(data)

			let data = ` + stringJSON + `;

			$$('datatable_part_1').clearAll();
			$$('datatable_part_2').clearAll();
			$$('datatable_part_3').clearAll();

			$$('datatable_part_1').parse(data);
			$$('datatable_part_2').parse(data);
			$$('datatable_part_3').parse(data);

		`

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
			fmt.Println("Catch error::Open file::OpenData", err)
			return
		}
	}
}
