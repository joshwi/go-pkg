package utils

import (
	"path"
	"regexp"
	"strings"
)

var a0 = regexp.MustCompile(`[^a-zA-Z\d\/\-]+`)
var a1 = regexp.MustCompile(`\_\-|\-\_`)
var a2 = regexp.MustCompile(`\_{2,}`)

func FormatPath(filepath string) string {

	ext := path.Ext(filepath)

	dest := strings.ReplaceAll(filepath, ext, "")

	dest = a0.ReplaceAllString(dest, "_")
	dest = a1.ReplaceAllString(dest, "-")
	dest = a2.ReplaceAllString(dest, "_")

	path := dest + ext

	return path
}
