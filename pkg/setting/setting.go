package setting

import "fmt"

const AppName = "user"

var (
	GitInfo    string
	AppVersion string
)

func init() {
	GitInfo = fmt.Sprintf("%s %s", GitBranch, GitVersion)
	AppVersion = fmt.Sprintf("%s built by %s at %s with %s", AppName, AppBuilder, AppBuildTime, GitInfo)
}
