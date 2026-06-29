package version

import (
	"fmt"
	"runtime"
)

var (
	Version   = "dev"
	Commit    = "unknown"
	Branch    = "unknown"
	BuildTime = "unknown"
)

type Info struct {
	Version   string
	Commit    string
	Branch    string
	BuildTime string
	GoVersion string
	GOOS      string
	GOARCH    string
}

func Get() Info {
	return Info{
		Version:   Version,
		Commit:    Commit,
		Branch:    Branch,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
		GOOS:      runtime.GOOS,
		GOARCH:    runtime.GOARCH,
	}
}

func (i Info) String() string {
	return fmt.Sprintf(
		"version=%s commit=%s branch=%s built=%s go=%s %s/%s",
		i.Version,
		i.Commit,
		i.Branch,
		i.BuildTime,
		i.GoVersion,
		i.GOOS,
		i.GOARCH,
	)
}
