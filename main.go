package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	settings *Settings
	ui       *UI
	isBreak  bool
	breakNum int
	sound    *Sound
)

func formatTime(tm time.Duration) string {
	s := int(tm.Seconds())
	minutes := int(s/60) % 60
	seconds := int(s % 60)

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func onMainTimerFinish() {
	isBreak = !isBreak
	newTime := settings.timer
	s := BackgroundColor
	e := BackgroundColorBreak

	// Update breaks guts.
	if isBreak {
		newTime = DefaultBreaks[breakNum]
		ui.breaks[breakNum].enable()
		breakNum = (breakNum + 1) % len(DefaultBreaks)
	} else {
		s = BackgroundColorBreak
		e = BackgroundColor
	}

	ui.bg.animate(s, e, BackgroundAnimationTime)

	// Update main timer.
	ui.timer.timer.stop()
	ui.timer.set(newTime)

	if settings.autoStartEnabled {
		ui.timer.timer.countDown()
	} else {
		ui.timer.started = false
	}

	// Dim all the break widgets when we cycle back to the beginning.
	if !isBreak && breakNum == 0 {
		ui.disableBreaks()
	}

	if settings.soundEnabled {
		sound.play(sound.cache["notification"])
	}
}

func main() {
	settings = &DefaultSettings

	sound = new(Sound)
	sound.init()
	sound.cache["notification"] = sound.open(NotificationSound)

	ui = new(UI)
	ui.app = app.New()

	ui.app.Settings().SetTheme(&newTheme{})
	ui.window = ui.app.NewWindow(WindowTitle)

	ui.createTray()

	ui.window.SetContent(ui.createContent())
	ui.window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	ui.window.SetMaster()
	ui.window.CenterOnScreen()
	ui.window.RequestFocus()
	ui.window.SetFixedSize(true)
	ui.window.SetCloseIntercept(func() { ui.window.Hide() })
	ui.window.ShowAndRun()
}
