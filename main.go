package main

import (
	"github.com/zserge/webview"
	"perv/engine/controller"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

type app struct {
	engine *controller.EngineController
}


func main() {
	app := app{}
	app.engine = new(controller.EngineController)
	url := app.engine.StartServer()
	w := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "Simple window demo",
		Resizable: false,
		URL:       url,
		ExternalInvokeCallback: app.engine.HandleRPC,
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	w.Run()
}