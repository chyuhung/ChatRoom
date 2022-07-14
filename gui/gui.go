package gui

import (
	"ChatRoom/client"
	"fmt"
	"fyne.io/fyne"
	app2 "fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func StartUI(c client.Client) {
	app := app2.New()
	loginWindow := app.NewWindow("登录")
	input := widget.NewEntry()
	input.Disabled()
	input.Resize(fyne.NewSize(24, 5))
	label := widget.NewLabel("Please input your name:")
	button := widget.NewButton("login", func() {
		if len(input.Text) > 0 {
			c.SetName(input.Text)
			label.Hidden = true

			input.SetText("")
			input.Hidden = true
			changeWindow(loginWindow, c)
		}
	})
	loginWindow.SetContent(widget.NewVBox(
		label,
		input,
		button,
	))
	loginWindow.Resize(fyne.NewSize(24, 24))
	loginWindow.ShowAndRun()
}

func changeWindow(window fyne.Window, c client.Client) {
	history := widget.NewMultiLineEntry()
	history.Disable()
	history.Resize(fyne.NewSize(480, 300))
	input := widget.NewEntry()
	//input.Disable()
	input.Resize(fyne.NewSize(460, 20))
	send := widget.NewButton("send", func() {
		if len(input.Text) > 0 {
			fmt.Println("Send start")
			c.SendMess(input.Text)
			input.SetText("")
		}
	})
	send.Resize(fyne.NewSize(20, 20))
	group := container.NewHBox(input, send)
	group.Resize(fyne.NewSize(480, 20))
	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), history, group)
	content.Resize(fyne.NewSize(480, 320))
	window.SetContent(content)
	window.Resize(fyne.NewSize(480, 320))

	go func() {
		for msg := range c.InComing() {
			AddMessage(history, msg.Name, msg.Message)
		}
	}()
}
func AddMessage(history *widget.Entry, user string, msg string) {
	history.SetText(history.Text + "\n" + user + ":" + msg)

}
