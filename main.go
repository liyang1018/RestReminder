package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne/config"
	"fyne/text"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"strconv"
	"time"
)

var (
	remainderWindow fyne.Window
	settingWindow   fyne.Window
	currentSecond   int
	c               *cron.Cron
	desk            desktop.App
	ok              bool
)

func main() {
	currentSecond = time.Now().Second() % 60
	initConfig()
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DefaultTheme())
	remainderWindow = myApp.NewWindow("Rest Remainder")
	settingWindow = myApp.NewWindow("Setting")
	remainderWindow.CenterOnScreen()
	settingWindow.CenterOnScreen()
	if desk, ok = myApp.(desktop.App); ok {
		desk.SetSystemTrayMenu(getMenuItem())
	}
	// content
	remainderWindow.SetContent(remainderContent(config.GlobalConfig.Time.RestDuration))
	settingWindow.SetContent(settingContent())
	// size
	remainderWindow.Resize(fyne.NewSize(float32(config.GlobalConfig.Window.ReminderWindowWidth), float32(config.GlobalConfig.Window.ReminderWindowHeight)))
	settingWindow.Resize(fyne.NewSize(float32(config.GlobalConfig.Window.SettingWindowWidth), float32(config.GlobalConfig.Window.SettingWindowHeight)))
	// close intercept
	remainderWindow.SetCloseIntercept(func() {
		remainderWindow.Hide()
	})
	settingWindow.SetCloseIntercept(func() {
		settingWindow.Hide()
	})
	icon, _ := fyne.LoadResourceFromPath("Icon.png")
	myApp.SetIcon(icon)

	myApp.Run()
}

func initConfig() {
	config.ViperConfig()
	text.InitText()
	cronjobToShowRemainderWindow()
}

func getMenuItem() *fyne.Menu {
	m := fyne.NewMenu("Rest Reminder",
		fyne.NewMenuItem("Next Rest Time: "+getNextRemainderTime(), nil),
		fyne.NewMenuItem("Skip To Next Rest", func() {
			currentSecond = time.Now().Second() % 60
			ShowRemainderWindow()
		}),
		fyne.NewMenuItem("Setting", func() {
			settingWindow.Show()
		}),
	)
	return m
}
func cronjobToShowRemainderWindow() {
	c = cron.New(cron.WithSeconds())
	cronExp := cast.ToString(currentSecond) + " 0/" + cast.ToString(config.GlobalConfig.Time.RestTime) + " * * * ?"
	println(cronExp)
	_, err := c.AddFunc(cronExp, func() {
		ShowRemainderWindow()
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func getNextRemainderTime() string {
	a := c.Entries()
	if len(a) == 0 {
		return "No remainder time"
	}

	return a[0].Next.Format("15:04")
}

func delayHideRemainderWindow(duration int) {
	time.Sleep(time.Duration(duration) * time.Minute)
	remainderWindow.Hide()
}

func ShowRemainderWindow() {
	remainderWindow.SetContent(remainderContent(config.GlobalConfig.Time.RestDuration))
	remainderWindow.Show()
	// not support dynamic refresh menu item
	desk.SetSystemTrayMenu(getMenuItem())
	delayHideRemainderWindow(config.GlobalConfig.Time.RestDuration)
}

func remainderContent(duration int) fyne.CanvasObject {
	label1 := widget.NewLabel(text.GetRandomRemindText())
	label1.TextStyle = fyne.TextStyle{Bold: true, Italic: true, Monospace: true, Symbol: true, TabWidth: 4}

	progressBar := widget.NewProgressBar()
	progressBar.Min = 0
	progressBar.Max = float64(config.GlobalConfig.Time.RestDuration * 60)

	button := widget.NewButton("Skip this break", func() {
		remainderWindow.Hide()
	})

	label2 := widget.NewLabel("")
	label2.Alignment = fyne.TextAlignCenter

	content := container.New(layout.NewVBoxLayout(),
		label1,
		progressBar,
		label2,
		button,
	)

	go func() {
		for i := duration*60 - 1; i >= 0; i-- {
			time.Sleep(time.Second)
			minutes := i / 60
			seconds := i % 60
			label2.SetText(
				strconv.Itoa(minutes) + " minutes " + strconv.Itoa(int(seconds)) + " seconds remaining",
			)
			progressBar.SetValue(float64(i))
		}
	}()

	return content
}

func settingContent() fyne.CanvasObject {
	return widget.NewLabel("Setting placeholder , rest-time, rest duration, window size, theme, language and so on")
}
