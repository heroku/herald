# Herald — Yielding Messages (for Languages) from Upon High

This application exists to empower Language Owners of new releases for their
respective languages and supporting libraries.

The notification system is comprised of a few different layers:

1. Scraping various "sources of truth" over the internet for release information (provided via buildpacks as standard Linux executables)
2. Presenting this information to the Language Owners
3. Notifying them, via GitHub Issues, when a new release has been made.

All languages, except Node.js, should benefit greatly from this system.

This system is written in the Go programming language.

## Relevant Links

- [heroku-herald GitHub profile](https://github.com/heroku-herald)

## Buildpack API

- Each builpack contains a `versions` directory, containing executables that print out, one per line, available version numbers.
- Each executable provided is a "target" that the buildpack needs to track versions of.

☤

Boostrapping Locally
--------------------

Install [dep](https://github.com/golang/dep):

    $ go get -u github.com/golang/dep/cmd/dep

Then, checkout this repository and resolve it's dependencies:

    $ dep ensure

That's it!