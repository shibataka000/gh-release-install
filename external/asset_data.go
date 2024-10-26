package external

var defaultAssetTemplates = map[string]AssetTemplateList{
	"hashicorp/terraform": {
		mustNewAssetTemplateFromString(""),
	},
}
