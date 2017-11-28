package herald

// import "os"
import "fmt"
// import "time"
import "github.com/hashicorp/go-getter"
import "io/ioutil"
import "log"
import "path/filepath"

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
	return fmt.Sprintf("<Buildpack name='%s'>", b.Name)
}

func (b Buildpack) FindVersionScripts() []string {
	log.Print(fmt.Sprintf("%s/versions/*", b.Path))
	results, _ := filepath.Glob(fmt.Sprintf("%s/versions/*", b.Path))
	return results
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

// func ExecutesBuildpacks() {

//  glob('*/updater/detect'
 
// }