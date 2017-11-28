package main

import "github.com/heroku/herald"
// todo: name that herlad
import "time"
import "log"



func main() {

	for {

		// Do the buldpack thing. 
		herald.DownloadBuildpacks()

		log.Print("Sleeping for 5 minutes…")
		
		// Sleep for five minutes. 
		time.Sleep(5*time.Minute)

		}

}