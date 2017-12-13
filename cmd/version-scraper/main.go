package main

import "github.com/heroku/herald"
import "github.com/fatih/color"
import "time"
import "log"
import "fmt"

func main() {

	// Redis stuff.
	redis := herald.NewRedis("")

	// Color Stuff.
	color.NoColor = false

	red := color.New(color.FgRed).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	bold := color.New(color.Bold, color.FgWhite).SprintFunc()

	for {

		// Get a list of the buildpacks (as types).
		buildpacks := herald.GetBuildpacks()

		// Iterate over them.
		for _, bp := range(buildpacks) {

			// Download and extract each Buildpack.
			log.Printf(bold("Downloading '%s'…"), red(bp.Name))
			path := bp.Download()

			log.Printf("Buildpack '%s' downloaded to '%s'!", red(bp), green(path))

			// Find all version executables for the given buildpack.
			executables := bp.FindVersionScripts()

			for _, exe := range(executables) {

				log.Printf(yellow("Executing '%s:%s' script…"), red(bp), magenta(exe))

				// TODO: Ensure chmod for the executable.
				exe.EnsureExecutable()

				// Execute the executable, print the results.
				results := exe.Execute()

				for _, result := range(results) {
					key := fmt.Sprintf("%s:%s:%s", bp, exe, result)
					value := "UNKNOWN"

					// Store the results in Redis.
					redis.Connection.Do("SETNX", key, value)


				}

				// Log results.
				log.Printf("%s:%s results: %s", red(bp), magenta(exe), results)

			}
		}

		log.Print(bold("Sleeping for 10 minutes…"))

		// Sleep for ten minutes.
		time.Sleep(10*time.Minute)

		}

}