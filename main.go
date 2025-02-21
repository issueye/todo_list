package main

import (
	"embed"
	"todo_list/server"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := server.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "待办事项",
		Width:         1200,
		Height:        800,
		DisableResize: false,
		Frameless:     true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 245, G: 245, B: 245, A: 1},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		Bind:             []interface{}{app},
		Windows: &windows.Options{
			WindowIsTranslucent:               true,
			WebviewIsTransparent:              true,
			BackdropType:                      windows.Mica,
			DisableWindowIcon:                 true,
			DisableFramelessWindowDecorations: true,
			WebviewUserDataPath:               "",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
