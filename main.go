package main

import (
	"embed"
	_ "embed"
	"log"
	"runtime"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/icons"
	"golang.design/x/hotkey"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

var (
	window *application.WebviewWindow
)

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "prism-go",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&GreetService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: false,
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	window = app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "Window 1",
		Frameless: true,
		Mac: application.MacWindow{
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
			WindowLevel:             application.MacWindowLevelFloating,
			InvisibleTitleBarHeight: 50,
		},
		URL:              "/",
		BackgroundColour: application.NewRGBA(0, 0, 0, 0),
		// BackgroundType:   application.BackgroundTypeTransparent,
		Width:         600,
		Height:        50,
		DisableResize: true,
		KeyBindings: map[string]func(window *application.WebviewWindow){
			"escape": func(window *application.WebviewWindow) {
				window.Hide()
			},
		},
	})

	systemTray := app.NewSystemTray()

	if runtime.GOOS == "darwin" {
		systemTray.SetTemplateIcon(icons.SystrayMacTemplate)

	} else {
		systemTray.SetDarkModeIcon(icons.SystrayDark)
		systemTray.SetIcon(icons.SystrayLight)
	}

	myMenu := app.NewMenu()
	myMenu.Add("Hello World!").OnClick(func(_ *application.Context) {
		app.NewWebviewWindowWithOptions((application.WebviewWindowOptions{
			Title: "Hello",
			URL:   "/#/page1",
		}))
	})
	systemTray.SetMenu(myMenu)

	window.OnWindowEvent(events.Common.WindowLostFocus, func(e *application.WindowEvent) {
		window.Hide()
	})

	go handleHotkey()
	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}

func handleHotkey() {
	showHideHotkey := hotkey.New([]hotkey.Modifier{hotkey.ModOption}, hotkey.KeySpace)
	if err := showHideHotkey.Register(); err != nil {
		log.Println(err)
		return
	}

	go func() {
		for range showHideHotkey.Keydown() {
			// log.Println("Pressed")
			if window.IsVisible() {
				window.Hide()
				log.Println(window.IsFocused())
			} else {
				window.Show()
				window.Focus()
				log.Println(window.IsFocused())
			}
		}
	}()
}
