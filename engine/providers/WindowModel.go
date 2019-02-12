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

func (m *WindowModel) IndexHTML() string {
	return `
<!doctype html>
<html>
	<head>
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
		
			webix.ui({
  				rows:[
  				    {
						cols:[
					    	{
					    	    id: 'resulted_text',
  				        		view: "textarea",
  				        		height: 400,
  				        		width:700,
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
        					{ id:"1",    header:"А",   width:24},
        					{ id:"2",   header:"Б",    width:24},
        					{ id:"3",    header:"В",   width:24},
        					{ id:"4",   header:"Г",    width:24},
        					{ id:"5",    header:"Д",   width:24},
        					{ id:"6",   header:"Е",    width:24},
        					{ id:"7",    header:"Ё",   width:24},
        					{ id:"8",   header:"Ж",    width:24},
        					{ id:"9",    header:"З",   width:24},
        					{ id:"10",   header:"И",   width:24},
        					{ id:"11",    header:"Й",  width:24},
        					{ id:"12",   header:"К",   width:24},
        					{ id:"13",    header:"Л",  width:24},
        					{ id:"14",   header:"М",   width:24},
        					{ id:"15",    header:"Н",  width:24},
        					{ id:"16",   header:"О",   width:24},
        					{ id:"17",    header:"П",  width:24},
        					{ id:"18",   header:"Р",   width:24},
        					{ id:"19",    header:"С",  width:24},
        					{ id:"20",   header:"Т",   width:24},
        					{ id:"21",    header:"У",  width:24},
        					{ id:"22",   header:"Ф",   width:24},
        					{ id:"23",    header:"Х",  width:24},
        					{ id:"24",   header:"Ц",   width:24},
        					{ id:"25",    header:"Ч",  width:24},
        					{ id:"26",   header:"Ш",   width:24},
        					{ id:"27",    header:"Щ",  width:24},
        					{ id:"28",   header:"Ъ",   width:24},
        					{ id:"29",    header:"Ы",  width:24},
        					{ id:"30",   header:"Ь",   width:24},
        					{ id:"31",    header:"Э",  width:24},
        					{ id:"32",   header:"Ю",   width:24},
        					{ id:"33",    header:"Я",  width:24},
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
	case dt == "close":
		wb.Terminate()
	case dt == "fullscreen":
		wb.SetFullscreen(true)
	case dt == "unfullscreen":
		wb.SetFullscreen(false)
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
		openData := string(b)
		fmt.Println(openData)

		// form JS
		jsString := `$$('resulted_text').define(value:"2");`
		fmt.Println(jsString)
		err = wb.Eval(jsString)
		if err != nil {
			fmt.Println("Catch error::Open file::OpenData", err)
			return
		}
	case dt == "opendir":
		log.Println("open", wb.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", ""))
	case dt == "save":
		log.Println("save", wb.Dialog(webview.DialogTypeSave, 0, "Save file", ""))
	case dt == "message":
		wb.Dialog(webview.DialogTypeAlert, 0, "Hello", "Hello, world!")
	case dt == "info":
		wb.Dialog(webview.DialogTypeAlert, webview.DialogFlagInfo, "Hello", "Hello, info!")
	case dt == "warning":
		wb.Dialog(webview.DialogTypeAlert, webview.DialogFlagWarning, "Hello", "Hello, warning!")
	case dt == "error":
		wb.Dialog(webview.DialogTypeAlert, webview.DialogFlagError, "Hello", "Hello, error!")
	case strings.HasPrefix(dt, "changeTitle:"):
		wb.SetTitle(strings.TrimPrefix(dt, "changeTitle:"))
	case strings.HasPrefix(dt, "changeColor:"):
		hex := strings.TrimPrefix(strings.TrimPrefix(dt, "changeColor:"), "#")
		num := len(hex) / 2
		if !(num == 3 || num == 4) {
			log.Println("Color must be RRGGBB or RRGGBBAA")
			return
		}
		i, err := strconv.ParseUint(hex, 16, 64)
		if err != nil {
			log.Println(err)
			return
		}
		if num == 3 {
			r := uint8((i >> 16) & 0xFF)
			g := uint8((i >> 8) & 0xFF)
			b := uint8(i & 0xFF)
			wb.SetColor(r, g, b, 255)
			return
		}
		if num == 4 {
			r := uint8((i >> 24) & 0xFF)
			g := uint8((i >> 16) & 0xFF)
			b := uint8((i >> 8) & 0xFF)
			a := uint8(i & 0xFF)
			wb.SetColor(r, g, b, a)
			return
		}
	}
}
