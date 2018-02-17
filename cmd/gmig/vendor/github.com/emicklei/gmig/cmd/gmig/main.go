package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

const version = "0.1"

var verbose = false

func main() {
	app := cli.NewApp()
	app.Version = version
	app.EnableBashCompletion = true
	app.Name = "gmig"
	app.Usage = "Google Cloud Platform infrastructure migration tool"

	// override -v
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "v",
			Usage: "verbose logging",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "new",
			Usage:  "gmig new \"create tester service account\"",
			Action: cmdCreateMigration,
		},
		{
			Name:        "up",
			Usage:       "gmig up",
			Description: "The up command runs the do section of all pending migrations in order, one after the other.",
			Action:      cmdMigrationsUp,
		},
		{
			Name:        "down",
			Usage:       "gmig down",
			Description: "The down command runs the undo section of the last applied migration only.",
			Action:      cmdMigrationsDown,
		},
		{
			Name:   "status",
			Usage:  "gmig status",
			Action: cmdMigrationsStatus,
		},
		{
			Name:   "init",
			Usage:  "gmig init",
			Action: cmdInit,
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
