package main

import (
	webview "github.com/webview/webview_go"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Block Display Engine")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://bdengine.app")
	w.Run()
}
