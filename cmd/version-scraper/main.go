package main

import "github.com/heroku/herald"
import "github.com/fatih/color"
import "github.com/mailgun/mailgun-go"
import "time"
import "log"
import "fmt"
import "os"

var MAILGUN_DOMAIN = os.Getenv("MAILGUN_DOMAIN")
var MAILGUN_API_KEY = os.Getenv("MAILGUN_API_KEY")
var MAILGUN_PUBLIC_KEY = os.Getenv("MAILGUN_PUBLIC_KEY")


func send_email(to string, version string) {
    
    mg := mailgun.NewMailgun(MAILGUN_DOMAIN, MAILGUN_API_KEY, MAILGUN_PUBLIC_KEY)
    message := mg.NewMessage(
        to,
        "Fancy subject!",
        "Hello from Mailgun Go!",
        "me@sandbox13d17507cb14496a9d97fe500600ac68.mailgun.org")
        resp, id, err := mg.Send(message)
        if err != nil {
            log.Fatal(err)
        }
   fmt.Printf("ID: %s Resp: %s\n", id, resp)
}


func main() {

	// Redis stuff.
	redis := herald.NewRedis("")

	// Color Stuff.
	color.NoColor = false

	red := color.New(color.FgRed).SprintFunc()
    blue := color.New(color.FgBlue).SprintFunc()
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
					value := herald.NewVersion().JSON()

					// Store the results in Redis.
					result, err := redis.Connection.Do("SETNX", key, value)

                    // The insert was successful (e.g. it didn't exist already)
                    if result.(int64) != int64(0) {
                        // TODO: Send a notification to the buildpack owner.
                        fmt.Println("Notifying %s about %s.", blue(bp.Owner), red(key))
//                         send_email(bp.Owner, key)
                    }
                    
					if err != nil {
						log.Fatal(err)
					}

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