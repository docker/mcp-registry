package licenses

import (
	"strings"

	"github.com/google/go-github/v70/github"
)

func IsValid(license *github.License) bool {
	if license != nil && (strings.HasPrefix(license.GetKey(), "gpl") || strings.HasPrefix(license.GetKey(), "agpl") || strings.HasPrefix(license.GetKey(), "npl")) {
		return false
	}
	return true
}
