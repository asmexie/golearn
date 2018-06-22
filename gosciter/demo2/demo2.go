package main

import (
	"log"

	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		&sciter.Rect{0, 0, 500, 500})
	if err != nil {
		log.Fatal(err)
	}
	w.LoadFile("demo2.html")
	w.SetTitle("表单")
	w.Show()
	w.Run()
}
