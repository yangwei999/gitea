package merlin

import (
	"bytes"
	"encoding/base64"
	"errors"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

type FrontMatter struct {
	License string `yaml:"license"`
}

var validLicense = map[string]bool{
	"apache-2.0":         true,
	"mit":                true,
	"cc-by-sa-3.0":       true,
	"afl-3.0":            true,
	"cc-by-sa-4.0":       true,
	"lgpl-3.0":           true,
	"lgpl-lr":            true,
	"cc-by-nc-3.0":       true,
	"bsd-2-clause":       true,
	"ecl-2.0":            true,
	"cc-by-nc-sa-4.0":    true,
	"cc-by-nc-4.0":       true,
	"gpl-3.0":            true,
	"cc0-1.0":            true,
	"cc":                 true,
	"bsd-3-clause":       true,
	"agpl-3.0":           true,
	"wtfpl":              true,
	"artistic-2.0":       true,
	"postgresql":         true,
	"gpl-2.0":            true,
	"isc":                true,
	"eupl-1.1":           true,
	"pddl":               true,
	"bsd-3-clause-clear": true,
	"mpl-2.0":            true,
	"odbl-1.0":           true,
	"cc-by-4.0":          true,
}

func CheckLicense(content string) error {
	b, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{}),
	)

	ctx := parser.NewContext()
	var buf bytes.Buffer
	if err = md.Convert(b, &buf, parser.WithContext(ctx)); err != nil {
		return err
	}

	data := frontmatter.Get(ctx)
	if data == nil {
		return errors.New("front matter is empty")
	}

	var meta FrontMatter
	if err = data.Decode(&meta); err != nil {
		return err
	}

	if _, ok := validLicense[meta.License]; !ok {
		return errors.New("invalid license")
	}

	return nil
}
