package version

import (
	"fmt"
	"runtime"
)

var (
	Version    = "0.0.1"
	Build      = ""
	StrVersion = fmt.Sprintf("scanner version %v, build %v %v", Version, Build, runtime.Version())
)
