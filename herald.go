package herald

// import "os"
import "fmt"
// import "time"
import "github.com/hashicorp/go-getter"
import "io/ioutil"
import "log"
import "path/filepath"
import "strings"

const BP_BRANCH = "versions"
const BP_TARBALL_TEMPLATE = "https://github.com/heroku/heroku-buildpack-%s/archive/%s.zip"
var BUILDPACKS = []string { "python", "php", "nodejs", "ruby", "jvm-common" }


type Version string

type Buildpack struct{
	Versions []Version
	Tarball string
	Path string
	// ExecutablePath string
	Name string
}

type Executable struct{
	Path string
}

func (e Executable) String() string {
	sl := strings.Split(e.Path, "/")
	return sl[len(sl) - 1]
}

func (b Buildpack) ZipballURI() string {
	return fmt.Sprintf(BP_TARBALL_TEMPLATE, b.Name, BP_BRANCH)
}

func (b Buildpack) BPDownload() (Buildpack, string) {
	target, _ := ioutil.TempDir("", "buildpacks")

	log.Printf("Downloading '%s'…", b.Name)

	getter.Get(target, b.ZipballURI())

	// The branch to base this off of. 

	// Set the Path.
	b.Path = target + fmt.Sprintf("/heroku-buildpack-%s-%s", b.Name, BP_BRANCH)

	return b, target
	
}

func (b Buildpack) String() string {
	return b.Name
}

func (b Buildpack) FindVersionScripts() []Executable {
	results := []Executable{}

	glob_results, _ := filepath.Glob(fmt.Sprintf("%s/versions/*", b.Path))
	for _, result := range(glob_results) {
		results = append(results, NewExecutable(result))
	}

	return results
}

func NewExecutable(path string) Executable {
	return Executable{
		Path: path,
	}
}

func NewBuildpack(name string) Buildpack {
	return Buildpack{
		Name: name,
	}
}

func GetBuildpacks() []Buildpack {
	// Download and unpack each Zipball from GitHub. 

	buildpacks := []Buildpack{}
	for _, bp := range(BUILDPACKS) {
		buildpacks = append(buildpacks, NewBuildpack(bp))
	}
	
	return buildpacks
}