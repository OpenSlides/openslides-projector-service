package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/leonelquinteros/gotext/cli/xgotext/parser"
	"github.com/leonelquinteros/gotext/cli/xgotext/parser/dir"
)

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// `some template {{Loc.Get "foo"}}` => some template bar
	reg := regexp.MustCompile(`\{\{\s*Loc\.Get?\s*["` + "`](.+?)[`" + `"][^}]*?\}\}`)

	data := &parser.DomainMap{
		Default: "default",
	}
	err := filepath.Walk("templates", func(path string, info fs.FileInfo, err error) error {
		noErr(err)
		if filepath.Ext(info.Name()) == ".html" && !info.IsDir() {
			b, err := os.ReadFile(path)
			noErr(err)
			sms := reg.FindAllStringSubmatch(string(b), -1)
			for _, s := range sms {
				data.AddTranslation("default", &parser.Translation{
					MsgID:           s[1],
					SourceLocations: []string{path},
				})
			}
		}
		return nil
	})
	noErr(err)

	err = dir.ParseDirRec("pkg", []string{}, data, false)
	noErr(err)

	err = data.Save("locale")
	noErr(err)
}
