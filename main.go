package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

var months = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

const (
	title               = "My Personal Diary ..."
	appWidth            = 900
	appHeight           = 600
	layoutOffset        = 0.25
	darkTheme           = "Dark"
	lightTheme          = "Light"
	themeButtonLabel    = "Change Theme"
	dropdownPlaceholder = "Select Month"
	titlePlaceholder    = "Title goes here ..."
	todayLabel          = "TODAY - "
	appId               = "github.harish.diary"
	noRecord            = "No memoir found ..."
)

func main() {
	a = app.NewWithID(appId)
	w := a.NewWindow(title)
	w.SetContent(LoadUI())
	w.Resize(fyne.NewSize(appWidth, appHeight))
	w.ShowAndRun()
}
