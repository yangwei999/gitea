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
	"apache-2.0": true,
	"mit":        true,
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
