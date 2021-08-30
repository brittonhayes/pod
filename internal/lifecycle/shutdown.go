package lifecycle

import "fyne.io/fyne/v2"

func Shutdown(app fyne.App, reason string, err error) {
	fyne.LogError(reason, err)
	app.Quit()
}
