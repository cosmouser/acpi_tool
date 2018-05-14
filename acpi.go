package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"os/exec"
	"time"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("ACPI tool")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)
	frame := gtk.NewFrame("test")
	vbox.PackStart(frame, false, false, 0)
	window.Add(vbox)
	window.ShowAll()
	go func() {
		for {
			out, err := exec.Command("acpi").Output()
			if err != nil {
				panic(err)
			}
			frame.SetLabel(string(out))
			time.Sleep(time.Minute * 1)
		}
	}()
	gtk.Main()

}
