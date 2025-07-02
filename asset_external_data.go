package main

import "text/template"

// defaultExternalAssetTemplates are templates of known release asset hosted on server other than GitHub.
var defaultExternalAssetTemplates = map[Repository][]ExternalAssetTemplate{
	Repository{Host: "github.com", Owner: "gravitational", Name: "teleport"}: {
		must(parseExternalAssetTemplate("https://cdn.teleport.dev/teleport-v{{.SemVer}}-linux-amd64-bin.tar.gz")),
	},
	Repository{Host: "github.com", Owner: "hashicorp", Name: "terraform"}: {
		must(parseExternalAssetTemplate("https://releases.hashicorp.com/terraform/{{.SemVer}}/terraform_{{.SemVer}}_linux_amd64.zip")),
	},
	Repository{Host: "github.com", Owner: "helm", Name: "helm"}: {
		must(parseExternalAssetTemplate("https://get.helm.sh/helm-{{.Tag}}-linux-amd64.tar.gz")),
	},
	Repository{Host: "github.com", Owner: "kubernetes", Name: "kubernetes"}: {
		must(parseExternalAssetTemplate("https://dl.k8s.io/release/{{.Tag}}/bin/linux/amd64/kubectl")),
	},
}

// parseExternalAssetTemplate returns a new [ExternalAssetTemplate] object.
func parseExternalAssetTemplate(downloadURL string) (ExternalAssetTemplate, error) {
	tmpl, err := template.New("DownloadURL").Parse(downloadURL)
	if err != nil {
		return ExternalAssetTemplate{}, err
	}
	return ExternalAssetTemplate{
		downloadURL: tmpl,
	}, nil
}
