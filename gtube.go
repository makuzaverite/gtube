package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/makuzaverite/gtube/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "gtube"
	app.Usage = "Get videos from youtube with your cli"

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "get current version of the app",
	}

	app.Version = "0.0.0"

	app.Authors = []*cli.Author{
		{
			Name:  "Makuza Mugabo Verite",
			Email: "mugaboverite@gmail.com",
		},
	}

	app.Action = func(c *cli.Context) error {

		if len(c.Args().Get(0)) <= 0 {
			fmt.Println("\nNo url specified")
			fmt.Print("gtube: use --help  for help\n")
		} else {
			cmd.DownloadDefaultVideo(c.Args().Get(0))
		}

		return nil
	}

	// flags
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "url",
			Aliases: []string{"u"},
			Usage:   "Video url from youtube",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	// start the CLI
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
