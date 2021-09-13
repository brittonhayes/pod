package config

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

// //go:embed logo.png
// var logo []byte

// WailsInit performs setup when Wails is ready.
func (c *Config) WailsInit(runtime *wails.Runtime) error {
	c.Runtime = runtime
	c.Logger = c.Runtime.Log.New("Config")

	// systray.Register(onReady, onExit)
	// defer go systray.Run(onReady, onExit)
	c.Logger.Info("Pod system tray started...")

	c.Logger.Info("Config initialized...")

	Run()
	c.Logger.Info("Directory setup complete...")
	return nil
}

// func onReady() {
// 	systray.SetIcon(logo)
// 	systray.SetTitle("pod")
// 	systray.SetTooltip("pro audio project utility")

// 	projects := systray.AddMenuItem("View Projects", "View projects in files")
// 	go tray.ViewProjects(projects)

// 	clients := systray.AddMenuItem("View Clients", "View clients in files")
// 	go tray.ViewClients(clients)

// 	systray.AddSeparator()

// 	quitBtn := systray.AddMenuItem("Quit", "Quit the application")
// 	go tray.Quit(quitBtn)
// }

// func onExit() {
// 	log.Info().Msg("exiting application")
// 	os.Exit(0)
// }
