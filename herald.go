package herald

// import "os"
import "fmt"
import "github.com/hashicorp/go-getter"
import "io/ioutil"
import "log"
import "path/filepath"
import "strings"
import "os"
import "os/exec"
import "encoding/json"
import "time"

// Buildpack Information.
const BP_BRANCH = "versions"
const BP_TARBALL_TEMPLATE = "https://github.com/heroku/heroku-buildpack-%s/archive/%s.zip"

// A buildpack that is owned by an owner (for notifications). 
type OwnedBuildpack struct {
    Name    string
    Owner   string
}

var BUILDPACKS = []OwnedBuildpack{
    {Name: "python", Owner: "kreitz@salesforce.com"}, 
    {Name: "php", Owner: "kreitz@salesforce.com"},
    {Name: "nodejs", Owner: "kreitz@salesforce.com"},
    {Name: "ruby", Owner: "kreitz@salesforce.com"},
    {Name: "jvm-common", Owner: "kreitz@salesforce.com"},
}

type Version struct {
	Name        string
	Target      Target
	Published   string `json:"id"`
	IsValid     bool   `json:"is_valid"`
	IsPublished bool   `json:"is_published"`
}

func NewVersion() Version {

	t := time.Now().UTC().Format(time.RFC3339)

	return Version{
		Published:   t,
		IsValid:     true,
		IsPublished: false,
	}
}

func (v Version) String() string {
	return fmt.Sprintf(
		"<Version name='%s' published=%#v, valid=%#v>",
		v.Name,
		v.IsPublished,
		v.IsValid,
	)
}

func (v Version) JSON() []byte {
	b, _ := json.Marshal(v)
	return b
}

// A Buildpack, which seems inherintly useful for this utility.
type Buildpack struct {
	Tarball string
	Path    string
	Name    string
    Owner   string
}

// Returns the GitHub ZipBall URI for the given buildpack.
func (b Buildpack) ZipballURI() string {
	return fmt.Sprintf(BP_TARBALL_TEMPLATE, b.Name, BP_BRANCH)
}

// Downloads the given buildpack to a temporary directory.
// Returns a new Buildpack object, as well as the target.
func (b *Buildpack) Download() string {
	target, _ := ioutil.TempDir("", "buildpacks")

	getter.Get(target, b.ZipballURI())

	// The branch to base this off of.

	// Set the Path.
	b.Path = target + fmt.Sprintf("/heroku-buildpack-%s-%s", b.Name, BP_BRANCH)

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

// Finds executables from a given buildpack, with globbing.
// Rerturns a slice of the Executable type.
func (b Buildpack) FindVersionScripts() []Executable {
	results := []Executable{}

	glob_results, _ := filepath.Glob(fmt.Sprintf("%s/versions/*", b.Path))
	for _, result := range glob_results {

		// Only yield a result if the glob result is a file.
		is_directory, _ := isDirectory(result)
		if !is_directory {
			results = append(results, NewExecutable(result))
		}

	}

	return results
}

// Creates a new Buildpack type.
func NewBuildpack(name string, owner string) Buildpack {
	return Buildpack{
		Name: name,
        Owner: owner,
	}
}

// Returns Targets for a given buildpack.
func (b Buildpack) GetTargets() []Target {
	redis := NewRedis(REDIS_URL)
	return redis.GetTargets(b.Name)
}

type Target struct {
	Buildpack Buildpack
	Name      string
	Versions  []Version
}

// Returns Versions for a given target.
func (t Target) GetVersions() []Version {
	redis := NewRedis(REDIS_URL)
	return redis.GetTargetVersions(t.Buildpack.Name, t.Name)
}

// Returns a Given Version for a given target.
func (t Target) GetVersion(version string) Version {
	redis := NewRedis(REDIS_URL)
	versions := redis.GetTargetVersions(t.Buildpack.Name, t.Name)
	new_version := Version{}

	for _, v := range versions {
		if v.Name == version {
			new_version = v
		}
	}

	return new_version
}

func NewTarget(bp Buildpack, name string) Target {

	return Target{
		Buildpack: bp,
		Name:      name,
		Versions:  nil,
	}
}

// An Executable, provided by a buildpack, for collecting version information.
type Executable struct {
	Path string
}

// String representation of Executable.
func (e Executable) String() string {
	sl := strings.Split(e.Path, "/")
	return sl[len(sl)-1]
}

// Ensures that the given executable is… executable.
func (e Executable) EnsureExecutable() {
	// TODO: Chmod to the proper permissions.
	if err := os.Chmod(e.Path, 0777); err != nil {
		// TODO: return error, etc.
		log.Fatal(err)
	}
}

// Executes the given executable, and returns results.
func (e Executable) Execute() []string {
	out, err := exec.Command(e.Path).Output()
	if err != nil {
		// TODO: Update this to return, etc.
		log.Fatal(err)
	}
	return strings.Split(strings.Trim(string(out), "\n"), "\n")

}

// Creates a new Executable type.
func NewExecutable(path string) Executable {
	return Executable{
		Path: path,
	}
}

// Generates a list of Buildpack objects, to be used by this utility.
func GetBuildpacks() []Buildpack {
	// Download and unpack each Zipball from GitHub.

	buildpacks := []Buildpack{}
	for _, bp := range BUILDPACKS {
		buildpacks = append(buildpacks, NewBuildpack(bp.Name, bp.Owner))
	}

	return buildpacks
}
