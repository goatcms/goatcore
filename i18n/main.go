package i18n

// I18N contains translation plain map
type I18N interface {
	// Set main translations
	Set(values map[string]string)
	// SetDefault default translations (can be overwritten by Set).
	SetDefault(values map[string]string)
	// Translate return translate for key
	Translate(key string, values ...interface{}) (string, error)
}
