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


// TODO: Maybe remove. 
type Version string


// A Buildpack, which seems inherintly useful for this utility. 
type Buildpack struct{
	Versions []Version
	Tarball string
	Path string
	// ExecutablePath string
	Name string
}

// An Executable, provided by a buildpack, for collecting version information. 
type Executable struct{
	Path string
}

// String representation of Executable. 
func (e Executable) String() string {
	sl := strings.Split(e.Path, "/")
	return sl[len(sl) - 1]
}

// Returns the GitHub ZipBall URI for the given buildpack. 
func (b Buildpack) ZipballURI() string {
	return fmt.Sprintf(BP_TARBALL_TEMPLATE, b.Name, BP_BRANCH)
}

// Downloads the given buildpack to a temporary directory.
// Returns a new Buildpack object, as well as the target. 
func (b Buildpack) BPDownload() (Buildpack, string) {
	target, _ := ioutil.TempDir("", "buildpacks")

	log.Printf("Downloading '%s'…", b.Name)

	getter.Get(target, b.ZipballURI())

	// The branch to base this off of. 

	// Set the Path.
	b.Path = target + fmt.Sprintf("/heroku-buildpack-%s-%s", b.Name, BP_BRANCH)

	return b, target
	
}

// String representatino of Buildpack. 
func (b Buildpack) String() string {
	return b.Name
}

// Finds executables from a given buildpack, with globbing.
// Rerturns a slice of the Executable type. 
func (b Buildpack) FindVersionScripts() []Executable {
	results := []Executable{}

	glob_results, _ := filepath.Glob(fmt.Sprintf("%s/versions/*", b.Path))
	for _, result := range(glob_results) {
		results = append(results, NewExecutable(result))
	}

	return results
}

// Creates a new Executable type. 
func NewExecutable(path string) Executable {
	return Executable{
		Path: path,
	}
}

// Creates a new Buildpack type. 
func NewBuildpack(name string) Buildpack {
	return Buildpack{
		Name: name,
	}
}

// Generates a list of Buildpack objects, to be used by this utility. 
func GetBuildpacks() []Buildpack {
	// Download and unpack each Zipball from GitHub. 

	buildpacks := []Buildpack{}
	for _, bp := range(BUILDPACKS) {
		buildpacks = append(buildpacks, NewBuildpack(bp))
	}
	
	return buildpacks
}