package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
)

var (
	timer_text binding.String
)

func on_tick(timer *Timer) {
	timer.show(timer_text)
}

func on_finish(timer *Timer) {
}

func main() {
	timer := create_timer(on_tick, on_finish)
	timer.set(TEST_TIMER)
	timer.start()

	ui := new(UI)
	ui.app = app.New()
	ui.timer = timer
	ui.window = ui.app.NewWindow(WINDOW_TITLE)
	ui.window.SetContent(ui.create_content())
	ui.window.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	ui.window.SetMaster()
	ui.window.CenterOnScreen()
	ui.window.RequestFocus()
	ui.window.SetFixedSize(true)
	ui.window.ShowAndRun()
}
