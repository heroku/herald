package lib

import "os"
import "fmt"
import "time"
import "github.com/hashicorp/go-getter"
import "io/ioutil"
import "log"

const BP_TARBALL_TEMPLATE = "https://github.com/heroku/heroku-buildpack-%s/archive/master.zip"
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
    target, err := ioutil.TempDir("", "buildpacks")
	if err != nil {
		log.Fatal(err)
	}
    
	
	// Download and unpack each Zipball from GitHub. 
    for _, tb := range(getBuildpackZipballs()) {
        getter.Get(target, tb)
    }
}

func ExecutesBuildpacks() {
 glob('*/updater/detect'
 
}