package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Little Command Line Interface"
	app.Usage = "Search IPs and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs by online adress",
			Flags:  flags,
			Action: searchIP,
		},
		{
			Name:   "servers",
			Usage:  "Search server names by online adress",
			Flags:  flags,
			Action: searchServerNames,
		},
	}

	return app
}

func searchIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServerNames(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range ips {
		fmt.Println(server.Host)
	}
}
