package providers

import (
	"fmt"
	"github.com/zserge/webview"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var GlobalData string

type WindowModel struct{}

func NewWindowModel() *WindowModel {
	return &WindowModel{}
}

func (m *WindowModel) IndexHTML() string {
	return `
<!doctype html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<link rel="stylesheet" href="http://cdn.webix.com/edge/webix.css" type="text/css"> 
    	<script src="http://cdn.webix.com/edge/webix.js" type="text/javascript"></script>  
	</head>
	<body>
		<!--<button onclick="external.invoke('open')">Open</button>-->
		<!--<button onclick="external.invoke('save')">Save</button>-->
		<!--<button onclick="external.invoke('changeTitle:'+document.getElementById('new-title').value)">-->
			<!--Analyze-->
		<!--</button>-->
		<!--<button onclick="external.invoke('changeTitle:'+document.getElementById('new-title').value)">-->
			<!--Analyze-->
		<!--</button>-->
		<!--<input id="new-title" type="text" />-->
		<script>
		
		let nullData = [[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]]
		JSON.stringify()
			webix.ui({
  				rows:[
  				    {
						cols:[
					    	{
					    	    id: 'resulted_text',
  				        		view: "textarea",
  				        		height: 400,
  				        		width:700,
  				        		on:{
    								'onChange': function(id){ 
            							// webix.message("Text loaded");
            						}
        						},
  				    		},
  				    		{
					    		rows:[
					        		{
    									view:"button", 
    									id:"open_button", 
    									value:"Open", 
    									type:"form", 
    									inputWidth:100,
    									on:{
    										'onItemClick': function(id){ 
            									external.invoke('open');
            								}
        								}
									},
									{
    									view:"button", 
    									id:"save_button", 
    									value:"Save", 
    									type:"form", 
    									inputWidth:100,
    									on:{
    										'onItemClick': function(id){ 
            									external.invoke('save');
            								}
        								}
									},
									{
    									view:"button", 
    									id:"exec_button", 
    									value:"Count", 
    									type:"form", 
    									inputWidth:100,
    									on:{
    										'onItemClick': function(id){
            									let text = $$('resulted_text').getValue();
            									external.invoke('push_table:' + text);
            								}
        								}
									},
					    		],
					    	},
						],
					},
      				{ 
      				    view:"template", 
        				type:"header", 
        				template:"Частотный анализ!",
        				tip: 'Составить таблицу'
        			},
      				{ 
      				    autowidth: true,
      				    view:"datatable",
      				    value: 0, 
        				columns:[
        					{ id:"1",    header:"А",   width:48},
        					{ id:"2",   header:"Б",    width:48},
        					{ id:"3",    header:"В",   width:48},
        					{ id:"4",   header:"Г",    width:48},
        					{ id:"5",    header:"Д",   width:48},
        					{ id:"6",   header:"Е",    width:48},
        					{ id:"7",    header:"Ё",   width:48},
        					{ id:"8",   header:"Ж",    width:48},
        					{ id:"9",    header:"З",   width:48},
        					{ id:"10",   header:"И",   width:48},
        					{ id:"11",    header:"Й",  width:48},
        					{ id:"12",   header:"К",   width:48},
        					{ id:"13",    header:"Л",  width:48},
        					{ id:"14",   header:"М",   width:48},
        					{ id:"15",    header:"Н",  width:48},
        					{ id:"16",   header:"О",   width:48},
        					{ id:"17",    header:"П",  width:48},
    					],
    					data: nullData,
    					scrollX: false,
    					scrollY: false,
      				},
					{ 
      				    autowidth: true,
      				    view:"datatable",
      				    value: 0, 
        				columns:[
        					{ id:"18",   header:"Р",   width:51},
        					{ id:"19",    header:"С",  width:51},
        					{ id:"20",   header:"Т",   width:51},
        					{ id:"21",    header:"У",  width:51},
        					{ id:"22",   header:"Ф",   width:51},
        					{ id:"23",    header:"Х",  width:51},
        					{ id:"24",   header:"Ц",   width:51},
        					{ id:"25",    header:"Ч",  width:51},
        					{ id:"26",   header:"Ш",   width:51},
        					{ id:"27",    header:"Щ",  width:51},
        					{ id:"28",   header:"Ъ",   width:51},
        					{ id:"29",    header:"Ы",  width:51},
        					{ id:"30",   header:"Ь",   width:51},
        					{ id:"31",    header:"Э",  width:51},
        					{ id:"32",   header:"Ю",   width:51},
        					{ id:"33",    header:"Я",  width:51},
    					],
    					data: nullData,
    					scrollX: false,
    					scrollY: false,
      				}
  				]
			});
		</script>
	</body>
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
		fmt.Println(strings.TrimPrefix(dt, "push_table:"))

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
