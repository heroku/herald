package main

import "github.com/heroku/herald"
// todo: name that herlad
import "time"
import "log"



func main() {

	for {

		// Get a list of the buildpacks (as types).
		buildpacks := herald.GetBuildpacks()

		// Iterate over them.
		for _, bp := range(buildpacks) {
			// Download and extract each Buildpack.
			path := bp.Download()

			log.Printf("Buildpack '%s' downloaded to '%s'", bp, path)
		}

		log.Print("Sleeping for 5 minutes…")

		// Sleep for five minutes. 
		time.Sleep(5*time.Minute)

		}

}