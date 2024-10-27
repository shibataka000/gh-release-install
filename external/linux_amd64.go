package external

import (
	"maps"

	"github.com/shibataka000/gh-release-install/github"
)

var defaultPatternForLinuxAmd64 = map[string]string{}

func init() {
	maps.Copy(DefaultPattern, github.DefaultPatterns)
	maps.Copy(DefaultPattern, defaultPatternForLinuxAmd64)
}
