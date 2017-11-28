package main

import "github.com/heroku/herald"
// todo: name that herlad
import "time"
import "log"



func main() {

	for {

		// Do the buldpack thing. 
		buildpacks := herald.GetBuildpacks()
		for _, bp := range(buildpacks) {
			log.Print(bp.Download())
		}

		log.Print("Sleeping for 5 minutes…")

		// Sleep for five minutes. 
		time.Sleep(5*time.Minute)

		}

}