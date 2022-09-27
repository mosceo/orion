package orion

import (
	"github.com/mosceo/mars"
)

func MarsVersion() string {
	return mars.GetVer()
}

func Version() string {
	return "1.8"
}
