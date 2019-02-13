package main

import (
	"fmt"
	"github.com/zserge/webview"
	"perv/engine/controller"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var GlobalData string

type app struct {
	logic  *controller.MainLogic
	engine *controller.EngineController
}

func main() {
	app := app{}
	app.engine = new(controller.EngineController)
	app.logic = new(controller.MainLogic)
	url, err := app.engine.StartServer()
	if err != nil {
		fmt.Println("Error while starting server::", err)
		return
	}
	w := webview.New(webview.Settings{
		Width:                  windowWidth,
		Height:                 windowHeight,
		Title:                  "Simple window demo",
		Resizable:              false,
		URL:                    url,
		ExternalInvokeCallback: app.engine.HandleRPC,
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	w.Run()
}
