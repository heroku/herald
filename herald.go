package herald

// import "os"
import "fmt"
// import "time"
import "github.com/hashicorp/go-getter"
import "io/ioutil"
import "log"

const BP_TARBALL_TEMPLATE = "https://github.com/heroku/heroku-buildpack-%s/archive/master.zip"
var BUILDPACKS = []string { "python", "php", "nodejs" }


type Version string

type Buildpack struct{
	Versions []Version
	Tarball string
	Path string
	// ExecutablePath string
	Name string
}

func (b Buildpack) ZipballURI() string {
	return fmt.Sprintf(BP_TARBALL_TEMPLATE, b.Name)
}

func (b Buildpack) Download() string {
	target, _ := ioutil.TempDir("", "buildpacks")

	log.Printf("Downloading '%s'…", b.Name)

	getter.Get(target, b.ZipballURI())

	return target
	
}

func NewBuildpack(name string) Buildpack {
	return Buildpack{
		// Path: nil,
		// ExecutablePath: path + "/bin/execute",
		Name: name,
	}
}

func GetBuildpacks() []Buildpack {
	// Download and unpack each Zipball from GitHub. 

	buildpacks := []Buildpack{}
	for _, bp := range(BUILDPACKS) {
	// log.Printf("Downloading '%s'…", bp)

	// err := getter.Get(target, tb)
	// if err != nil {
	// 	return err
		buildpacks = append(buildpacks, NewBuildpack(bp))
	}
	
	return buildpacks
}

// func ExecutesBuildpacks() {

//  glob('*/updater/detect'
 
// }