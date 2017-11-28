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
			bp, path := bp.BPDownload()

			log.Printf("Buildpack '%s' downloaded to '%s'!", bp, path)

			executables := bp.FindVersionScripts()
			for _, exe := range(executables) {
				log.Printf("Execututing '%s:%s' script…", bp, exe)
			}
		}

		log.Print("Sleeping for 5 minutes…")

		// Sleep for five minutes. 
		time.Sleep(5*time.Minute)

		}

}