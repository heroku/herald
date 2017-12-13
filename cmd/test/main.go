package main

import "github.com/heroku/herald"

func main() {

	python := herald.NewBuildpack("python")
	python.GetTargets()
}