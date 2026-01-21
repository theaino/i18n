package i18n

import (
	"fmt"
	"strings"
)

const missingFormat string = "!(MISSING: %s)"

type Locale struct {
	Code Code
	Messages map[string]any
}

func (l *Locale) T(path string, args ...any) string {
	message, ok := l.Value(path).(string)
	if !ok {
		return fmt.Sprintf(missingFormat, path)
	}
	return fmt.Sprintf(message, args...)
}

func (l *Locale) Has(path string) bool {
	_, ok := l.Value(path).(string)
	return ok
}

func (l *Locale) Value(path string) any {
	return walkDict(l.Messages, pathParts(path))
}

func pathParts(path string) []string {
	return strings.Split(path, ".")
}

func walkDict(dict map[string]any, path []string) any {
	if len(path) == 0 {
		return nil
	}
	if len(path) == 1 {
		return dict[path[0]]
	}
	return walkDict(dict[path[0]].(map[string]any), path[1:])
}
