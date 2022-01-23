package main

import (
	"github.com/gotk3/gotk3/gtk"
)

// Setup the Window.
func setupWindow() *gtk.Window {
	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.Connect("destroy", gtk.MainQuit)
	w.SetDefaultSize(500, 300)
	w.SetPosition(gtk.WIN_POS_CENTER)
	w.SetTitle("TextView properties example")

	return w
}

func main() {
	gtk.Init(nil)

	win := setupWindow()
	box, _ := gtk.ListBoxNew()
	win.Add(box)

	btn, _ := gtk.ButtonNewWithLabel("test")
	
	box.Add(btn)

	win.ShowAll()

	gtk.Main()
}
