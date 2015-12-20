package main

import (
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
)

func main() {
    gtk.Init(nil)
    window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    window.SetPosition(gtk.WIN_POS_CENTER)
    window.SetTitle("Hello GTK+Go world!")
    window.SetIconName("gtk-dialog-info")
    window.Connect("destroy", func(ctx *glib.CallbackContext) {
        gtk.MainQuit()
    }, "foo")
    vbox := gtk.NewVBox(false, 1)
    button := gtk.NewButtonWithLabel("Hello world!")
    button.SetTooltipMarkup("Say Hello World to everybody!")
    button.Clicked(func() {
        messagedialog := gtk.NewMessageDialog(
            button.GetTopLevelAsWindow(),
            gtk.DIALOG_MODAL,
            gtk.MESSAGE_INFO,
            gtk.BUTTONS_OK,
            "Hello!")
        messagedialog.Response(func() {
            messagedialog.Destroy()
        })
        messagedialog.Run()
    })
    vbox.PackStart(button, false, false, 0)
    window.Add(vbox)
    window.SetSizeRequest(300, 50)
    window.ShowAll()
    gtk.Main()
}
