package app

import "github.com/urfave/cli"

func Build() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "Server fetcher"
	app.Usage = "Fetche ipv4 addresses and dns servers"
	return
}
