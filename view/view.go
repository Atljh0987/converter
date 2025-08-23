package view

import (
	"converter/global"
	"converter/page"
)

func RunWindow() {
	window := global.InitAppContainer().Window
	
	// window.SetContent(page.NewHomePage(window))
	window.SetContent(page.NewSettingsPage(window))

	window.ShowAndRun()
}
