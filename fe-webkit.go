package main

import (
	"io/ioutil"
	"runtime"

	"github.com/garekkream/BookShelf/Settings"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
)

func webkitInit() {
	runtime.LockOSThread()
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		Settings.Log().Fatal("Failed to create GTK Window!!")
	}

	win.SetTitle("Bookshelf")
	win.SetSizeRequest(1000, 800)
	win.SetResizable(false)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	glib.IdleAdd(func() bool {
		file, _ := ioutil.ReadFile("Web/index.html")

		webView.LoadHTML(string(file), "Web")
		return false
	})

	win.Add(webView)
	win.ShowAll()

	gtk.Main()
}
