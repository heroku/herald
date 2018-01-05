package main

import (
	"fmt"
	"github.com/heroku/herald"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()

	// --is-valid.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "is-invalid",
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

	// --is-published.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "unpublished",
			Usage: "Lists unpublished versions.",
		},
	}

	app.Action = func(c *cli.Context) error {
		// The Buildpack.

		parsed := strings.Split(c.Args().Get(0), ":")

		bp := parsed[0]

		// "python" was speficied.
		if len(parsed) == 1 {
			// If no buildpack was passed in, fail hard:
			if parsed[0] == "" {
				fmt.Println("A buildpack or target must be provided.")
				os.Exit(1)
			}

			// List the given buildpack's targets:
			fmt.Printf("Buildpack '%s' targets:\n\n", bp)

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

		targetString := parsed[1]

		// Get the buildpack from Redis.
		buildpack := herald.NewBuildpack(bp)

		// "python:python" was speficied.
		if len(parsed) == 2 {
			fmt.Printf("Buildpack '%s:%s' releases:\n\n", bp, targetString)

			// Get the targets for the given buildpack.
			target := herald.NewTarget(buildpack, targetString)

			versions := target.GetVersions()

			// Print them to the screen.
			for _, v := range versions {

				// if --unpublished passed, only display unpublished versions.
				if c.Bool("unpublished") {
					if !v.IsPublished {
						fmt.Printf(" - %s\n", v)
					}
				} else {
					fmt.Printf(" - %s\n", v)
				}

			}

			// Exit, because we're finished.
			os.Exit(0)
		}

		target := herald.NewTarget(buildpack, targetString)
		version := target.GetVersion(targetString)
		// "python:python:3.6.3" was speficied.
		fmt.Printf("Info: %s:%s:%s\n", bp, target.Name, targetString)

		// Get Version from given version information.
		fmt.Printf("  Valid: %t\n", version.IsValid)
		fmt.Printf("  Published: %t\n", version.IsPublished)

		// python := herald.NewBuildpack("python")
		return nil
	}

	app.Run(os.Args)
}
