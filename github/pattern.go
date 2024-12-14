package github

import (
	"bytes"
	"regexp"
	"slices"
	"strconv"
	"text/template"
)

// Pattern represents a pair of regular expression of GitHub release asset download URL and template of executable binary name.
// This is used to select an appropriate one from GitHub release assets and determine an executable binary name.
type Pattern struct {
	// asset is a regular expression of GitHub release asset download URL.
	// This is used to select an appropriate one from GitHub release assets and used as input data to determine an executable binary name.
	asset *regexp.Regexp

	// execBinary is a template of executable binary name.
	// This is used to determine an executable binary name.
	execBinary *template.Template
}

// newPattern returns a new [Pattern] object.
func newPattern(asset *regexp.Regexp, execBinary *template.Template) Pattern {
	return Pattern{
		asset:      asset,
		execBinary: execBinary,
	}
}

// newPatternFromString returns a new [Pattern] object.
func newPatternFromString(asset string, execBinary string) (Pattern, error) {
	a, err := regexp.Compile(asset)
	if err != nil {
		return Pattern{}, err
	}

	b, err := template.New("ExecBinary").Parse(execBinary)
	if err != nil {
		return Pattern{}, err
	}

	return newPattern(a, b), nil
}

// newPatternArrayFromStringMap returns a new array of [Pattern] objects.
// Map's keys should be regular expressions of GitHub release asset download URL and values should be templates of executable binary name.
func newPatternArrayFromStringMap(patterns map[string]string) ([]Pattern, error) {
	ps := []Pattern{}
	for asset, execBinary := range patterns {
		p, err := newPatternFromString(asset, execBinary)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

// match returns true if regular expression in pattern matches given GitHub release asset download URL.
func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.DownloadURL.String()))
}

// priority returns a literal prefix length of regular expression of GitHub release asset download URL as priority of pattern.
// Pattern with higher priority is prioritized over pattern with lower priority.
func (p Pattern) priority() int {
	prefix, _ := p.asset.LiteralPrefix()
	return len(prefix)
}

// execute applies a template of executable binary name to values of capturing groups in regular expression of GitHub release asset download URL and returns [ExecBinary] object.
func (p Pattern) execute(asset Asset) (ExecBinary, error) {
	data := map[string]string{}
	submatch := p.asset.FindStringSubmatch(asset.DownloadURL.String())

	for i := range submatch {
		data[strconv.Itoa(i)] = submatch[i]
	}

	for _, name := range p.asset.SubexpNames() {
		if name == "" {
			continue
		}
		index := p.asset.SubexpIndex(name)
		if index < 0 && index >= len(submatch) {
			return ExecBinary{}, errOutOfIndex
		}
		data[name] = submatch[index]
	}

	var b bytes.Buffer
	if err := p.execBinary.Execute(&b, data); err != nil {
		return ExecBinary{}, err
	}

	return newExecBinary(b.String()), nil
}

// findAssetAndPattern find [Asset] and [Pattern] matching and returns them.
// Pattern with higher priority is prioritized over pattern with lower priority.
func findAssetAndPattern(assets []Asset, patterns []Pattern) (Asset, Pattern, error) {
	cloned := slices.Clone(patterns)
	slices.SortFunc(cloned, func(p1, p2 Pattern) int {
		return p2.priority() - p1.priority()
	})

	for _, p := range cloned {
		for _, a := range assets {
			if p.match(a) {
				return a, p, nil
			}
		}
	}

	return Asset{}, Pattern{}, ErrNoAssetsMatchPattern
}
