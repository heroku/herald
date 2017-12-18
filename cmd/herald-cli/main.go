package main

import (
	"fmt"
	"github.com/heroku/herald"
	"github.com/urfave/cli"
	"os"
	"strings"
)

// func main() {

// 	python := herald.NewBuildpack("python")
// 	t := python.GetTargets()[0]
// 	fmt.Println(t.GetVersions())
// }

func main() {
	app := cli.NewApp()

	// --is-valid.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "is-valid",
			Usage: "Sets the given release as valid.",
		},
	}

	// --is-published.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "is-published",
			Usage: "Sets the given release as published.",
		},
	}

	app.Action = func(c *cli.Context) error {
		// The Buildpack.

		parsed := strings.Split(c.Args().Get(0), ":")

		bp := parsed[0]

		if len(parsed) == 1 {
			fmt.Printf("Buildpack %q targets:\n\n", bp)

			// Get the buildpack from Redis.
			buildpack := herald.NewBuildpack(bp)

			// Get the targets for the given buildpack.
			targets := buildpack.GetTargets()

			// Print them to the screen.
			for _, t := range targets {
				fmt.Printf(" - %q\n", t.Name)
			}

			os.Exit(0)
		}

		target := parsed[1]

		if len(parsed) == 2 {
			fmt.Printf("Buildpack %q:%q targets:\n\n", bp, target)

			// Get the buildpack from Redis.
			buildpack := herald.NewBuildpack(bp)

			// Get the targets for the given buildpack.
			target := herald.NewTarget(buildpack, target)

			versions := target.GetVersions()

			// Print them to the screen.
			for _, v := range versions {
				fmt.Printf(" - %q\n", v)
			}

			os.Exit(0)
		}

		release := parsed[2]

		fmt.Printf("Info: %q:%q:%s\n", bp, target, release)

		// python := herald.NewBuildpack("python")
		return nil
	}

	app.Run(os.Args)
}
