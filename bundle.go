package i18n

import (
	"io"
	"io/fs"
	"maps"

	"github.com/goccy/go-yaml"
)

type Bundle struct {
	Locales map[Code]*Locale
}

func LoadFS(bundleFS fs.FS) (*Bundle, error) {
	bundle := &Bundle{Locales: make(map[Code]*Locale)}
	err := fs.WalkDir(bundleFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		file, err := bundleFS.Open(path)
		if err != nil {
			return err
		}
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		var bundleData map[string]any
		err = yaml.Unmarshal(data, &bundleData)
		if err != nil {
			return err
		}
		bundle.Merge(LoadBundle(bundleData))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return bundle, nil
}

func LoadBundle(data map[string]any) *Bundle {
	bundle := &Bundle{
		Locales: make(map[Code]*Locale),
	}
	for code, messages := range data {
		bundle.Locales[Code(code)] = &Locale{
			Code: Code(code),
			Messages: messages.(map[string]any),
		}
	}
	return bundle
}

func (b *Bundle) Merge(other *Bundle) {
	maps.Copy(b.Locales, other.Locales)
}

func (b *Bundle) AddLocale(locale *Locale) {
	b.Locales[locale.Code] = locale
}

func (b *Bundle) GetLocale(code Code) *Locale {
	for matchingCode := range code.Matches() {
		if locale, ok := b.Locales[matchingCode]; ok {
			return locale
		}
	}
	return nil
}
