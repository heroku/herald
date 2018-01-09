# Herald — Yielding Messages (for Languages) from Upon High

**This repo is a work in progress.**

This application exists to empower Language Owners of new releases for their
respective languages and supporting libraries.

The notification system is comprised of a few different layers:

1. Scraping various "sources of truth" over the internet for release information (provided via buildpacks as standard Linux executables)
2. Presenting this information to the Language Owners
3. Notifying them, via GitHub Issues, when a new release has been made.

All languages, except Node.js, should benefit greatly from this system.

This system will be written in the Go programming language.

☤

Boostrapping Locally
--------------------

Install [dep](https://github.com/golang/dep):

    $ go get -u github.com/golang/dep/cmd/dep

Then, checkout this repository and resolve it's dependencies:

    $ dep ensure

That's it!