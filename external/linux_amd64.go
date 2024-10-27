package external

import (
	"maps"

	"github.com/shibataka000/gh-release-install/github"
)

var defaultPatternForLinuxAmd64 = map[string]string{}

func init() {
	maps.Copy(DefaultPatterns, github.DefaultPatterns)
	maps.Copy(DefaultPatterns, defaultPatternForLinuxAmd64)
}
