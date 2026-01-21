package i18n

import "context"

var LocaleKey = struct{}{}

func T(ctx context.Context, path string, args ...any) string {
	return GetLocale(ctx).T(path, args...)
}

func Has(ctx context.Context, path string) bool {
	return GetLocale(ctx).Has(path)
}

func Value(ctx context.Context, path string) any {
	return GetLocale(ctx).Value(path)
}

func GetLocale(ctx context.Context) *Locale {
	locale, ok := ctx.Value(LocaleKey).(*Locale)
	if !ok {
		return nil
	}
	return locale
}

func WithLocale(ctx context.Context, locale *Locale) context.Context {
	return context.WithValue(ctx, LocaleKey, locale)
}
