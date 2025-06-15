package github2

import (
	"bytes"
	"regexp"
	"slices"
	"strconv"
	"text/template"
)

// Pattern represents a pair of regular expression and template for asset selection and binary naming.
type Pattern struct {
	asset      *regexp.Regexp
	execBinary *template.Template
}

func newPattern(asset *regexp.Regexp, execBinary *template.Template) Pattern {
	return Pattern{
		asset:      asset,
		execBinary: execBinary,
	}
}

func parsePattern(asset string, execBinary string) (Pattern, error) {
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

func parsePatternMap(patterns map[string]string) ([]Pattern, error) {
	ps := []Pattern{}
	for asset, execBinary := range patterns {
		p, err := parsePattern(asset, execBinary)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

func (p Pattern) match(asset Asset) bool {
	return p.asset.Match([]byte(asset.DownloadURL.String()))
}

func (p Pattern) priority() int {
	prefix, _ := p.asset.LiteralPrefix()
	return len(prefix)
}

// Execute applies the template to the asset and returns ExecBinary.
func (p Pattern) Execute(asset Asset) (ExecBinary, error) {
	return p.execute(asset)
}

func (p Pattern) execute(asset Asset) (ExecBinary, error) {
	data := map[string]string{}
	submatch := p.asset.FindStringSubmatch(asset.DownloadURL.String())

	for i := range submatch {
		data[strconv.Itoa(i)] = submatch[i]
	}

	for _, name := range p.asset.SubexpNames() {
		index := p.asset.SubexpIndex(name)
		if index >= 0 && index < len(submatch) {
			data[name] = submatch[index]
		}
	}

	var b bytes.Buffer
	if err := p.execBinary.Execute(&b, data); err != nil {
		return ExecBinary{}, err
	}

	return newExecBinary(b.String()), nil
}

// FindResult represents the result of ApplicationService.Find.
type FindResult struct {
	Asset      Asset
	ExecBinary ExecBinary
}

// NewRelease returns a new Release object.
func NewRelease(tag string) Release {
	return newRelease(tag)
}

// ParsePatternMap returns a new array of Pattern objects from a map.
func ParsePatternMap(patterns map[string]string) ([]Pattern, error) {
	return parsePatternMap(patterns)
}

// FindAssetAndPattern finds Asset and Pattern matching and returns them.
func FindAssetAndPattern(assets []Asset, patterns []Pattern) (Asset, Pattern, error) {
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

// NewFindResult returns a new FindResult object.
func NewFindResult(asset Asset, execBinary ExecBinary) FindResult {
	return FindResult{
		Asset:      asset,
		ExecBinary: execBinary,
	}
}
