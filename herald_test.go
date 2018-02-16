package herald

import (
	// "fmt"
	"github.com/hashicorp/go-getter"
	// "os"
	"io/ioutil"
	"testing"
)

func TestBPExecution(t *testing.T) {
	target, _ := ioutil.TempDir("", "buildpacks")
	getter.Get(target, "https://github.com/heroku-herald/heroku-buildpack-testing/archive/master.zip")

	path := target + "/heroku-buildpack-testing-master"

	bp := Buildpack{
		Tarball: "https://github.com/heroku-herald/heroku-buildpack-testing/archive/master.zip",
		Path:    path,
		Name:    "Dummy",
		Owner:   "me@kennethreitz.org",
	}

	executables := bp.FindVersionScripts()
	if len(executables) != 1 {
		t.Errorf("Executables found was incorrect, got: %d, want: %d.", len(executables), 1)
	}

	for _, exe := range executables {
		exe.EnsureExecutable()
		results, _ := exe.Execute()

		if len(results) != 2 {
			t.Errorf("Results found was incorrect, got: %d, want: %d.", len(results), 2)
		}
	}
}
