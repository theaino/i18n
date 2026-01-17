package i18n

import (
	"iter"
	"strings"
)

type Code string

func (c Code) Matches() iter.Seq[Code] {
	return func(yield func(Code) bool) {
		codeParts := strings.Split(string(c), "-")
		for length := len(codeParts); length > 0; length-- {
			if !yield(Code(strings.Join(codeParts[:length], "-"))) {
				return
			}
		}
	}
}

func (c Code) Base() Code {
	return Code(strings.SplitN(string(c), "-", 2)[0])
}

func ParseAcceptLanguage(value string) []Code {
	codes := make([]Code, 0)
	for s := range strings.SplitSeq(value, ",") {
		s = strings.TrimSpace(s)

		if i := strings.Index(s, ";"); i > 0 {
			s = s[:i]
		}

		codes = append(codes, Code(s))
	}
	return codes
}
