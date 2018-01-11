package herald

import (
	"fmt"
	"github.com/hashicorp/go-getter"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// BPBranch is which branch of the buildpack to use.
const BPBranch = "versions"

// BPTarballTemplate is the template for constructing the zipball URL.
const BPTarballTemplate = "https://github.com/heroku/heroku-buildpack-%s/archive/%s.zip"

// OwnedBuildpack is a buildpack that is owned by an owner (for notifications).
type OwnedBuildpack struct {
	Name  string
	Owner string
}

// Buildpacks is a list of a buildpacks with their respective owners.
var Buildpacks = []OwnedBuildpack{
	{Name: "python", Owner: "kennethreitz"},
	{Name: "php", Owner: "kreitz@salesforce.com"},
	{Name: "nodejs", Owner: "kreitz@salesforce.com"},
	{Name: "ruby", Owner: "kreitz@salesforce.com"},
	{Name: "jvm-common", Owner: "kreitz@salesforce.com"},
}

// Version is a type used to represent a given version of a Target.
type Version struct {
	Name   string
	Target Target
}

// NewVersion returns a new Version instance.
func NewVersion() Version {
	return Version{}
}

// Buildpack is a type which seems inherintly useful for this utility.
type Buildpack struct {
	Tarball string
	Path    string
	Name    string
	Owner   string
}

// ZipballURI Returns the GitHub ZipBall URI for the given buildpack.
func (b Buildpack) ZipballURI() string {
	return fmt.Sprintf(BPTarballTemplate, b.Name, BPBranch)
}

// Download Downloads the given buildpack to a temporary directory.
// Returns a new Buildpack object, as well as the target.
func (b *Buildpack) Download() string {
	target, _ := ioutil.TempDir("", "buildpacks")

	getter.Get(target, b.ZipballURI())

	// The branch to base this off of.

	// Set the Path.
	b.Path = target + fmt.Sprintf("/heroku-buildpack-%s-%s", b.Name, BPBranch)

	return b.Path

}

// String representatino of Buildpack.
func (b Buildpack) String() string {
	return b.Name
}

// Determines wether a given path is a file or not.
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}

// FindVersionScripts Finds executables from a given buildpack, with globbing.
// Rerturns a slice of the Executable type.
func (b Buildpack) FindVersionScripts() []Executable {
	results := []Executable{}

	globResults, _ := filepath.Glob(fmt.Sprintf("%s/versions/*", b.Path))
	for _, result := range globResults {

		// Only yield a result if the glob result is a file.
		isDirectory, _ := isDirectory(result)
		if !isDirectory {
			results = append(results, NewExecutable(result))
		}

	}

	return results
}

// NewBuildpack Creates a new Buildpack type.
func NewBuildpack(name string, owner string) Buildpack {
	return Buildpack{
		Name:  name,
		Owner: owner,
	}
}

// Target represents a target (e.g. buildable asset) owned by a Buildpack.
type Target struct {
	Buildpack Buildpack
	Name      string
	Versions  []Version
}

// NewTarget returns a new Target instance.
func NewTarget(bp Buildpack, name string) Target {

	return Target{
		Buildpack: bp,
		Name:      name,
		Versions:  nil,
	}
}

// Executable provided by a buildpack, for collecting version information.
type Executable struct {
	Path string
}

// String representation of Executable.
func (e Executable) String() string {
	sl := strings.Split(e.Path, "/")
	return sl[len(sl)-1]
}

// EnsureExecutable Ensures that the given executable is… executable.
func (e Executable) EnsureExecutable() {
	// TODO: Chmod to the proper permissions.
	if err := os.Chmod(e.Path, 0777); err != nil {
		// TODO: return error, etc.
		log.Fatal(err)
	}
}

// Execute Executes the given executable, and returns results.
func (e Executable) Execute() []string {
	out, err := exec.Command(e.Path).Output()
	if err != nil {
		// TODO: Update this to return, etc.
		log.Fatal(err)
	}
	return strings.Split(strings.Trim(string(out), "\n"), "\n")

}

// NewExecutable Creates a new Executable type.
func NewExecutable(path string) Executable {
	return Executable{
		Path: path,
	}
}

// GetBuildpacks Generates a list of Buildpack objects, to be used by this utility.
func GetBuildpacks() []Buildpack {
	// Download and unpack each Zipball from GitHub.

	buildpacks := []Buildpack{}
	for _, bp := range Buildpacks {
		buildpacks = append(buildpacks, NewBuildpack(bp.Name, bp.Owner))
	}

	return buildpacks
}
