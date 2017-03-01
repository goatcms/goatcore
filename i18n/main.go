package i18n

type I18N interface {
	Set(values map[string]string)
	SetDefault(values map[string]string)

	Translate(key string, values ...interface{}) (string, error)
}
