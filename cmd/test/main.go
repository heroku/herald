package main

import "github.com/heroku/herald"
import "fmt"

func main() {

	python := herald.NewBuildpack("python")
	t := python.GetTargets()[0]
	fmt.Println(t.GetVersions())
}