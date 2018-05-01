# Herald — Yielding Messages (for Languages) from Upon High

![](https://travis-ci.com/heroku/herald.svg?token=7FY3Nqwfz5mTRqxuNHxx&branch=master)

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

## Add a Check to Herald

Herald assumes that you have a buildpack in a git repo that contains a `versions` branch, and in that versions branch there is a `versions` folder. In that folder it will look for executables. For example in the Ruby buildpack https://github.com/heroku/heroku-buildpack-ruby/tree/versions/versions.

When the executable is run, it should print out an available version number. One per line. For example in the Ruby buildpack:

```term
$ versions/ruby
2.5.1
2.4.4
2.3.7
2.2.10
2.6.0-preview1
#...
```

The list of buildpacks is in a go "map" like a dict in Python or a hash in Ruby in `herald.go`.

Once every 15 minutes the `herald/cmd/version-scraper/main.go` script will execute against all buildpacks and record the result in Redis. When it is detected that there is a new entry, then a new issue will be opened up on that buildpack.

☤

Boostrapping Locally
--------------------

Install [dep](https://github.com/golang/dep):

    $ go get -u github.com/golang/dep/cmd/dep

Then, checkout this repository and resolve it's dependencies:

    $ dep ensure

That's it!
