package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/garyburd/redigo/redis"
	"github.com/google/go-github/github"
	"github.com/heroku/herald"
	"github.com/stvp/rollbar"
	"golang.org/x/oauth2"
	"log"
	"os"
	"time"
)

// MinutesToSleep is the number of minutes the process sleeps after successfully
// running all language scripts.
var MinutesToSleep = 15

// GithubToken is Personal GitHub token.
var GithubToken = os.Getenv("GITHUB_TOKEN")

// RollbarToken is provided by Heroku Addon.
var RollbarToken = os.Getenv("ROLLBAR_ACCESS_TOKEN")

// Opens an issue on GitHub for the given buildpack and new target.
//
// Note: Uses the GITHUB_TOKEN environment variable, which is currently
//   Set to Kenneth's personal GitHub account. Need to create a bot account
//   for this service.
func openIssue(bp herald.Buildpack, target string) bool {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	title := fmt.Sprintf("New release (%s) available! (Herald System)", target)
	body := fmt.Sprintf("This issue created programmatically and automatically by Heroku, on behalf of %s, the owner of the %s buildpack.", bp.Owner, bp.Name)

	newIssue := github.IssueRequest{
		Title: &title,
		Body:  &body,
		// Labels: ["New Build Target"],
		Assignee: &bp.Owner,
	}
	// list all repositories for the authenticated user
	bpName := fmt.Sprintf("heroku-buildpack-%s", bp.Name)

	issue, _, err := client.Issues.Create(ctx, "heroku", bpName, &newIssue)
	if err != nil {
		fmt.Println("An error occurred creating the GitHub issue. Will try again")
		rollbar.Error(rollbar.ERR, err)
	} else {
		fmt.Println(fmt.Sprintf("New issue created on %s buildpack on GitHub.", bp.Name))
	}

	return (issue != nil)
}

func main() {

	// Redis stuff.
	Redis := herald.NewRedis("")

	// Rollbar stuff.
	rollbar.Token = RollbarToken
	rollbar.Environment = "production"

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
		for _, bp := range buildpacks {

			// Download and extract each Buildpack.
			log.Printf(bold("Downloading '%s'…"), red(bp.Name))
			path := bp.Download()

			log.Printf("Buildpack '%s' downloaded to '%s'!", red(bp), green(path))

			// Find all version executables for the given buildpack.
			executables := bp.FindVersionScripts()

			for _, exe := range executables {

				log.Printf(yellow("Executing '%s:%s' script…"), red(bp), magenta(exe))

				targetQuery, _ := redis.Strings(Redis.Connection.Do("KEYS", fmt.Sprintf("%s:%s:*", bp, exe)))
				targetCount := len(targetQuery)
				newTarget := (targetCount == 0)

				// Ensure chmod for the executable.
				exe.EnsureExecutable()

				// Execute the executable, print the results.
				results, err := exe.Execute()
				if err != nil {
					rollbar.Error(rollbar.ERR, err)
				}

				for _, result := range results {
					key := fmt.Sprintf("%s:%s:%s", bp, exe, result)
					value := 1

					// Store the results in Redis.
					result, err := Redis.Connection.Do("SETNX", key, value)
					if err != nil {
						rollbar.Error(rollbar.ERR, err)
					}

					// The insert was successful (e.g. it didn't exist already)
					if result.(int64) != int64(0) {
						// TODO: Send a notification to the buildpack owner.

						// Open an issue on GitHub (work in progress).
						if !newTarget {
							fmt.Println("Notifying", blue(bp.Owner), "about", red(key), ".")
							success := openIssue(bp, key)
							if !success {
								// If writing out the issue was unsuccessful, delete the key from Redis.
								_, err := Redis.Connection.Do("DEL", key)
								if err != nil {
									rollbar.Error(rollbar.ERR, err)
								}
							}
						} else {
							fmt.Println("New target, skipping GitHub notifications.")
						}

					}

					if err != nil {
						rollbar.Error(rollbar.ERR, err)
						log.Fatal(err)
					}

				}

				// Log results.
				log.Printf("%s:%s results: %s", red(bp), magenta(exe), results)
			}
		}

		log.Print(bold(fmt.Sprintf("Sleeping for %v minutes…", MinutesToSleep)))

		// Sleep for ten minutes.
		time.Sleep(time.Duration(MinutesToSleep) * time.Minute)

	}

}
