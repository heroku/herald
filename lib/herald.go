package lib

import "os"
import "fmt"
import "time"
import "github.com/hashicorp/go-getter"

var BP_TARBALL_TEMPLATE = "https://github.com/heroku/heroku-buildpack-%s/archive/master.zip"
var BUILDPACKS = []string { "python", "php", "nodejs" }

func getBuildpackZipballs() []string {
	collected := []string{}
	
	for _, bp := range(BUILDPACKS) {
		// Add comment here. 
		collected = append(collected, fmt.Sprintf(BP_TARBALL_TEMPLATE, bp))
	}
	
	return collected

}

func DownloadBuildpacks() {
	
	// create temp directory
	// use 'fake' for now
    target := "fake"
	
	// Download and unpack each Zipball from GitHub. 
    for _, tb := range(getBuildpackZipballs()) {
        getter.Get(target, tb)
    }
}

func ExecutesBuildpacks() {
 glob('*/updater/detect'
 
}

func add(a int, b int) int {
    fmt.Println("hello world")
    return a + b
}
