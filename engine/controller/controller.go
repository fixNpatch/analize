package controller

import (
	"github.com/zserge/webview"
	"log"
	"net"
	"net/http"
	"perv/engine/providers"
)

type EngineController struct {
	model *providers.WindowModel
	logic *providers.MainLogic
}

func (c *EngineController) Init() *EngineController {
	c.model = providers.NewWindowModel()
	return nil
}

func (c *EngineController) StartServer() (url string, err error) {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(c.model.IndexHTML()))
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	url = "http://" + ln.Addr().String()
	return url, nil
}

func (c *EngineController) HandleRPC(w webview.WebView, data string) {
	c.model.HandleRPC(&w, &data)
	return
}
