package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Build() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "Server fetcher"
	app.Usage = "Fetche ipv4 addresses and dns servers"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Fetch ipv4 address",
			Flags:  flags,
			Action: ip,
		},
		{
			Name:   "servers",
			Usage:  "Look for server names",
			Flags:  flags,
			Action: server,
		},
	}

	return
}

func ip(context *cli.Context) {
	ips, err := net.LookupIP(context.String("host"))
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func server(context *cli.Context) {
	servers, err := net.LookupNS(context.String("host"))
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server)
	}
}
