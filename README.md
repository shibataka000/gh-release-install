# gh-release-install

[![Test](https://github.com/shibataka000/gh-release-install/actions/workflows/test.yaml/badge.svg)](https://github.com/shibataka000/gh-release-install/actions/workflows/test.yaml)

Install executable binary from GitHub release asset.

## Usage

```
Install executable binary from GitHub release asset.

Usage:
  gh-release-install [flags]

Flags:
  -D, --dir string               Directory where executable binary will be installed into. (default ".")
  -h, --help                     help for gh-release-install
      --pattern stringToString   Map whose key should be regular expressions of GitHub release asset download URL to download and value should be templates of executable binary name to install. (default [(?i)^.+/(?P<name>[^\.]+)([\-\._]v?\d+\.\d+\.\d+)?[\-\._]linux([\-\._](amd64|x86_64|64bit))?(\.tar\.gz|\.zip|\.gz|\.tgz)?$={{.name}},https://github\.com/istio/istio/releases/download/.+/istioctl-\d+\.\d+\.\d+-linux-amd64\.tar\.gz$=istioctl,https://github\.com/starship/starship/releases/download/.+/starship-x86_64-unknown-linux-gnu\.tar\.gz$=starship,https://dl\.k8s\.io/release/.+/bin/linux/amd64/kubectl=kubectl,https://github\.com/protocolbuffers/protobuf/releases/download/.+/protoc-\d+\.\d+-linux-x86_64\.zip$=protoc])
  -R, --repo string              GitHub repository name. This should be [HOST/]OWNER/REPO format.
      --tag string               GitHub release tag.
```

## Install

```
go install github.com/shibataka000/gh-release-install@main
```
