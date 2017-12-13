package main

import "github.com/heroku/herald"
import "fmt"

func main() {

	python := herald.NewBuildpack("python")
	fmt.Println(python.GetTargets())
}