package i18mem

import "testing"

const (
	PLTranslate = "pl.form.validator.min_length"
	ENTranslate = "en.form.validator.min_length"
)

func TestSimpleLoad(t *testing.T) {
	t.Parallel()
	i18 := NewI18N()
	i18.Set(map[string]string{
		PLTranslate: "Minimalna długość pola to %v znaków",
		ENTranslate: "The minimum length of the field is %v characters",
	})
	i18.SetDefault(map[string]string{
		PLTranslate: "Zbyt krótki",
		ENTranslate: "Too short",
	})
	pl, err := i18.Translate(PLTranslate, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if "Minimalna długość pola to 2 znaków" != pl {
		t.Errorf("incorrect translation for key %s (%s)", PLTranslate, pl)
		return
	}
	en, err := i18.Translate(ENTranslate, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if "The minimum length of the field is 2 characters" != en {
		t.Errorf("incorrect translation for key %s (%s)", ENTranslate, en)
		return
	}
}
