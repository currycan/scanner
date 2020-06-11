package version

import (
	"fmt"
	"runtime"
)

var (
	Version    = "0.0.1"
	Build      = ""
	VersionStr = fmt.Sprintf("scanner version %v, build %v %v", Version, Build, runtime.Version())
)
