package main

import (
	"github.com/zserge/lorca"
	"log"
	"net/http"
)

func main() {
	ui, err := lorca.New("", "./browser-resources", 1280, 720)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	/*
	ui.Bind("Analyze", func(url string) string {
		log.Println(url)
		return "done"
	})
	 */

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", http.StripPrefix("/", fs))
	go http.ListenAndServe("127.0.0.1:3800", nil)
	// Load HTML after Go functions are bound to JS
	_ = ui.Load("http://127.0.0.1:3800")
	_ = ui.SetBounds(lorca.Bounds{Left: 0, Top: 0, WindowState: lorca.WindowStateMaximized})

	<-ui.Done()
}