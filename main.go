package main

import (
	"os"

	"github.com/DBosley/go-react-chat/server"
	"github.com/codegangsta/cli"
)

func main() {
	Run(os.Args)
}

// Run is the react app entry point
func Run(args []string) {
	app := cli.NewApp()
	app.Name = "go-chat"
	app.Usage = "React Chat"

	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "Run Server",
			Action: RunServer,
		},
	}
	app.Run(args)
}

// RunServer is the server entry point
func RunServer(c *cli.Context) {
	server.Start()
}
