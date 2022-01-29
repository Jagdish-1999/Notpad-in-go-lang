package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 0

func main() {
	a := app.New()
	w := a.NewWindow("")
	w.Resize(fyne.NewSize(800, 650))
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapBreak
	input.SetPlaceHolder("Enter Text Here")

	item2 := fyne.NewMenuItem("Open", func() {
		openDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)
				output := fyne.NewStaticResource("New File ", readData)
				viewData := widget.NewMultiLineEntry()
				viewData.Wrapping = fyne.TextWrapBreak
				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(string(output.StaticName))
				w.SetContent(container.NewScroll(viewData))
				w.Resize(fyne.NewSize(800, 650))
				w.Show()
			}, w)
		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openDialog.Show()
	})
	item3 := fyne.NewMenuItem("Save", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				uc.Write(textData)
			}, w)

		saveFileDialog.SetFileName("New File" + strconv.Itoa(count) + ".txt")
		saveFileDialog.Show()
	})
	item1 := fyne.NewMenuItem("New", func() {
		count++
		str := "New File" + strconv.Itoa(count)
		w.SetTitle("Notpad" + "/" + str)

	})

	menu1 := fyne.NewMenu("File", item1, item2, item3)
	menu2 := fyne.NewMenu("Edit")
	menu3 := fyne.NewMenu("View")
	menu4 := fyne.NewMenu("Help")

	mainMenu := fyne.NewMainMenu(menu1, menu2, menu3, menu4)
	w.SetMainMenu(mainMenu)

	w.SetContent(input)

	w.ShowAndRun()
}
